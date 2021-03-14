# Go开发Web应用（5）使用https

前言：数据安全是Web服务中必须考虑的重要因素，基于http的web应用或者接口，在提供服务时，传输时都会将信息以明文暴露，增加了信息窃取的风险。所以更加理想的方式是将服务升级为https的方式提供。接下来的学习内容便是如何在Go中使用https。

## 启动https

> [https是基于HTTP协议，通过SSL或TLS提供加密处理数据、验证对方身份以及数据完整性保护](https://blog.csdn.net/xiaoming100001/article/details/81109617)

为了安全起见，最好是使用`https`方式来访问web服务，在go中启动`https`方式十分简单。之前启动web服务的方式类似下面这样

```go
package main

import "net/http"

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello"))
	})

	_ = http.ListenAndServe("localhost:8080", nil)
}
```

访问时使用的是http://localhost:8080 ，如果需要将其改为 `https` 的方式只需要修改启动方法如下

````go
//_ = http.ListenAndServe("localhost:8080", nil)
//使用https
_ = http.ListenAndServeTLS("localhost:8080", "cert.pem", "key.pem", nil)
````

其中`"cert.pem"`和`"key.pem"`分别是证书和秘钥文件（私钥）的路径。如果没有现成的证书，在开发阶段可以使用自签证书来先用。生成自签证书可以使用openssl工具，但是go贴心的为我们提供了现成的方法。

运行`%goroot%\src\crypto\tls\generate_cert.go`文件即可生成证书，打开文件可以看到可用的命令行参数。

```go
var (
	host       = flag.String("host", "", "Comma-separated hostnames and IPs to generate a certificate for")
	validFrom  = flag.String("start-date", "", "Creation date formatted as Jan 1 15:04:05 2011")
	validFor   = flag.Duration("duration", 365*24*time.Hour, "Duration that certificate is valid for")
	isCA       = flag.Bool("ca", false, "whether this cert should be its own Certificate Authority")
	rsaBits    = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if --ecdsa-curve is set")
	ecdsaCurve = flag.String("ecdsa-curve", "", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	ed25519Key = flag.Bool("ed25519", false, "Generate an Ed25519 key")
)
```

使用默认参数，设置host即可生成可用的自签证书。打开命令行，切到当前工作目录，执行。

```bash
$ go run %goroot%\src\crypto\tls\generate_cert.go --host localhost

wrote cert.pem
wrote key.pem
```

成功后`"cert.pem"`和`"key.pem"`文件便生成在当前工作目录中。现在启动浏览器通过 https://localhost:8080 进行访问。但是会出现下面的画面，这是由于校验证书失败了，必须忽略这个校验才能得到响应。

<img src="https://raw.githubusercontent.com/Xuhy0826/Go-Go-Study/master/resource/https1.png" alt="访问https" style="zoom:50%;" />

同样的，使用postman或者curl方式访问，都会遇到类似的问题，都必须显式的跳过证书验证才能得到信息。那么使用go创建客户端来访问会如何呢？下面是客户端的示例代码。

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	r, err := http.Get("https://localhost:8080")
	if err != nil {
		log.Fatal("error:", err)
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	log.Println(string(body))
}
```

上面使用`net/http`包的Get方法发起请求，不出意外也是出现了同样的验证错误。

![go客户端访问](https://raw.githubusercontent.com/Xuhy0826/Go-Go-Study/master/resource/https2.png)

为了显式的跳过验证，客户端代码需要做以下改动

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, //跳过验证
			},
		},
	}
	//发起GET请求
	r, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	log.Println(string(body))
}
```

再次运行，现在可以正常得到响应。

![go客户端请求https](https://raw.githubusercontent.com/Xuhy0826/Go-Go-Study/master/resource/https3.png)

必须显式的跳过验证，是不是感觉用了个假的https，那么接下来就来解决这个问题。我们需要在客户端将证书加入到证书池中，如下示例

```go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//使用证书池，将自己签发的证书加进去
	pool := x509.NewCertPool()
	myCaPath := "../cert.pem"
	caCrt, _ := ioutil.ReadFile(myCaPath)
	pool.AppendCertsFromPEM(caCrt)

	//创建httpClient，并指定证书池
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            pool,
				InsecureSkipVerify: false, //这次不跳过验证
			},
		},
	}
	//发起GET请求
	r, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	log.Println(string(body))
}
```

现在再运行客户端，便可以正常得到响应了。除了使用go自带的生成证书的方式，我们也可以使用openssl来生成证书。

## 使用openssl自签发证书

接下来我们需要使用openssl来制作自签证书，Windows安装包版本可以在 [这里](http://slproweb.com/products/Win32OpenSSL.html) 找到对应的安装包进行下载。mac的话自带了openssl，所以就不需要在额外安装了。可以使用`openssl version`检查安装是否成功。

有了openssl之后，我在当前工作目录下创建了新的路径命名为`cert`，切进去依次运行下面的命令（如果想了解关于证书验证背后的原理 [这里](https://blog.csdn.net/xiaoming100001/article/details/81109617/) 介绍很详细）。

1. 生成CA的私钥
```bash
$ openssl genrsa -out ca.key 2048
```

2. 根据CA自己的私钥生成自签发的数字证书，该证书里包含CA自己的公钥。
```bash
$ openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost"  -days 5000 -out ca.crt
```

3. 生成证书请求
```bash
$ openssl req -new -sha256 -key ca.key \
-subj "/C=CN/ST=Macao/L=Macao/O=Kwh/OU=Itd/CN=localhost" \
-reqexts SAN \
-config <(cat /System/Library/OpenSSL/openssl.cnf \
        <(printf "[SAN]\nsubjectAltName=DNS:localhost,DNS:www.xuhy.top")) \
-out xuhy.csr
```

4. 签名证书
```bash
$ openssl x509 -req -days 365000 \
-in xuhy.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
-extfile <(printf "subjectAltName=DNS:localhost,DNS:www.xuhy.top") \
-out xuhy.crt
```

执行完成之后，会生成5个文件，根据这5个文件再修改下客户端和服务端的代码。

./server/main.go

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello"))
	})

	http.HandleFunc("/reg", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			data, _ := ioutil.ReadAll(request.Body)
			log.Printf("%v", string(data))
		}else {
			writer.WriteHeader(http.StatusNotFound)
		}
	})

	//_ = http.ListenAndServe("localhost:8080", nil)
	//使用https
	err := http.ListenAndServeTLS("localhost:8080", "../cert/xuhy.crt", "../cert/ca.key", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
```

启动服务端，成功后再修改客户端访问的代码，这时要将我们自签发的证书加进去。

./client/main.go

```go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//使用证书池，将自己签发的证书加进去
	pool := x509.NewCertPool()
	myCaPath := "../cert/ca.crt"
	caCrt, _ := ioutil.ReadFile(myCaPath)
	pool.AppendCertsFromPEM(caCrt)

	//创建httpClient，并指定证书池
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            pool,
				InsecureSkipVerify: false, //这次不跳过验证
			},
		},
	}
	//发起GET请求
	r, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	log.Println(string(body))

	//发起POST请求
	r1, err := client.Post("https://localhost:8080/reg", "text/plain", strings.NewReader("hello, i'm client"))
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	log.Println(r1.StatusCode)
}
```

现在在执行客户端访问，便可以成功得到结果了。

![客户端访问https](https://raw.githubusercontent.com/Xuhy0826/Go-Go-Study/master/resource/https4.png)

另外，如果是自己有域名的情况下，现在可以在 [Let's Encrypt](https://letsencrypt.org/zh-cn/) 上免费申请证书了。这部分内容我等我的域名备案通过之后再补充上来吧。当然处理直接使用https之外，也可以通过nginx进行反向代理的方式，这种方法就不在本文中讨论了。

#### 补充

值得注意的一点，我们将服务的启动方式更改为`ListenAndServeTLS`之后，可以在源码中看到，会默认使用http2进行通信。

```go
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error {
	// Setup HTTP/2 before srv.Serve, to initialize srv.TLSConfig
	// before we clone it and create the TLS Listener.
	if err := srv.setupHTTP2_ServeTLS(); err != nil {
		return err
	}
	
	... ...
	
}
```

使用http2可以大大提高其通信的效率，下一节要学习的grpc也是优先采用http2的方式通信的。