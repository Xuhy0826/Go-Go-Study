# 35

## ssl

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

运行``%goroot%\src\crypto\tls\generate_cert.go`文件即可生成证书，打开文件可以看到可用的命令行参数。

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

成功后`"cert.pem"`和`"key.pem"`文件便生成在当前工作目录中。现在启动即可通过 https://localhost:8080 进行访问。

