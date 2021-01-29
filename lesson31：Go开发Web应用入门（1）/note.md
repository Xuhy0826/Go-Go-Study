# Golang搭建Web应用入门（1）

### 开门见山
启动一个最简单的Web Server的示例。
```
package main

import "net/http"

func main() {
	startDefaultServer()
}
func startDefaultServer() {
	//第一个参数：路由地址
	//第二个参数：处理请求的函数
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("hello world"))
	})

	//开启http请求监听
	//参数1 ：地址，默认为 *:80
	//参数2 ：Handler接口类型，指定请求的处理器，默认是 DefaultServeMux
	http.ListenAndServe("localhost:8080", nil)
}
```
运行后访问 http://localhost:8080 即可看到响应。  
`http.ListenAndServe("localhost:8080", nil)`实际上是创建一个`http.Server`并调用其`ListenAndServe()`函数，此时便可启动监听。方法的第一个入参指定网络地址，第二个入参是指定`Server`的`Handler`，也就是请求的处理器。处理器能够来定义如何处理接收的请求。如果这个参数传递的是`nil`，那么就会使用一个默认的`Handler`，即`DefaultServeMux`。
> `http.Server` 结构定义
```
type Server struct {
    //指定网络地址，为空则默认为“*:80"
	Addr string

    //接口类型，用来处理请求，默认为http.DefaultServeMux
	Handler Handler
}
```
> `ListenAndServe()`函数声明，内容先不细究
```
func (srv *Server) ListenAndServe() error {
	....
}
```
启动监听后，当收到请求，会使用`Server`的`Handler`接口的`ServeHTTP`方法来处理请求。
`Handler`是一个接口，定义如下
```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```
`Handler`中的`ServeHTTP`方法，入参两个：
* `ResponseWriter`，接口，用于写响应
* `Request`指针，请求数据
> `ResponseWriter`接口定义
```
type ResponseWriter interface {
    Header() Header

    Write([]byte) (int, error)

    WriteHeader(statusCode int)
}
```
> `Request`结构的定义，未列全
```
type Request struct {
	Host 			string
	Method 			string
	URL 			*url.URL
	Header 			Header			//类型 map[string][]string
	Body 			io.ReadCloser

	Proto      		string 
	ProtoMajor 		int
	ProtoMinor 		int
	
	ContentLength 	int64			//body内容的长度
	Form 			url.Values
	PostForm 		url.Values
	...
}
```
通过`Request`这个结构，就可以获取本次请求的很多信息，这些属性名都是一目了然的。  

现在，了解这些之后，自定义一个`Server`和`Handler`来监听并处理请求。
```
package main

import "net/http"

func main() {
	//startDefaultServer()

	startMyServer()
}

func startMyServer() {
	//自定义server
	server := http.Server{
		Addr:    "localhost: 8080",
		Handler: myHandler{}, //不指定，默认为 DefaultServeMux
	}
	//启动监听
	server.ListenAndServe()
}

// 自定义 Handler
type myHandler struct{}

// myHandler 实现 Handler 接口
// 接收请求，返回"Hello gopher"字符串
func (h myHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello gopher"))
}
```

### DefaultServeMux 是什么
之前提到如果在创建`http.Server`时没有指定其`Handler`字段的值或赋`nil`值，那么便会使用`DefaultServeMux`作为`Handler`来处理请求。`DefaultServeMux`是一个默认的`ServeMux`，`ServeMux`是一个结构体，其定义如下。官方命名其为**多路复用器**，并且`DefaultServeMux`也实现 `Handler` 接口的。
> `DefaultServeMux`的定义
```
var DefaultServeMux = &defaultServeMux
var defaultServeMux ServeMux

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}
```
然而`DefaultServeMux`的正确打开方式是用来将对server的请求分发到不同的`Handler`的路由。我们首先知道`Handler`接口是用来处理请求，针对不同的请求地址应该制定不同的`Handler`进行处理，`DefaultServeMux`来作为前置的`Handler`来分发请求，所以相当于一个路由器，这也是“多路复用器”的含义。  
![DefaultServeMux路由请求到各个Handler](https://github.com/Xuhy0826/Golang-Study/blob/master/resource/httpHandler.png)

### 配置多个Handler
当`Server`的`Handler`使用默认的`DefaultServeMux`时（即Handler字段不赋值或赋nil）。使用`http.Handle()`函数便可将自定义的`Handler`“注册”到`DefaultServeMux`，这样访问不同的url就可以使用不同的`Handler`来处理。  
> `http.Handle()`函数的定义
```
// 入参：
// pattern：地址后跟的路由
// handler：对应pattern的Handler
func Handle(pattern string, handler Handler) { 
	DefaultServeMux.Handle(pattern, handler) 
}
```
可以看出，`http.Handle()`即是调用`DefaultServeMux`的`Handle()`方法。现通过`http.Handle()` 来向`DefaultServeMux`注册多个`Handler`。
先定义两个`Handler`。
```
type aHandler struct{}

func (ah aHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello gopher from aHandler"))
}

type bHandler struct{}

func (bh bHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello gopher from bHandler"))
}
```
定义Server，Handler使用`DefaultServeMux`，随后将刚刚定义的两个`Handler`注册到`DefaultServeMux`并分配相应的路由地址。
```
package main

import "net/http"

func main() {
	multiHandlerServer()
}

func multiHandlerServer() {
	server := http.Server{
		Addr:    "localhost: 8080",
		Handler: nil, //此时为 DefaultServeMux
	}
	http.Handle("/a", aHandler{})
	http.Handle("/b", bHandler{})

	server.ListenAndServe()
}
```
现启动Server后，访问 http://localhost:8080/a 便会返回`Hello gopher from aHandler`，http://localhost:8080/b 便会返回`Hello gopher from bHandler`。  
当然，也可以不用`DefaultServeMux`自己创建一个`ServeMux`类型。可以使用http包提供的`NewServeMux()`函数。
```
package main

import (
	"net/http"
	"time"
)

func main() {
	//创建 ServeMux
	mux := http.NewServeMux()

	mux.Handle("/", aHandler{})

	server := &http.Server{
		Addr:    "localhost: 8080",
		Handler: mux,
	}
	
	server.ListenAndServe()
}
```
### 再回首
如果现在回到开篇的第一个示例，可以发现这里用的并不是刚刚说的`http.Handle`方法，而是`http.HandleFunc`方法，这个方法的第一个参数和`http.Handle`方法一致，但是第二个参数是一个函数类型的参数。
```
package main

import "net/http"

func main() {
	startDefaultServer()
}

func startDefaultServer() {
	http.HandleFunc("/",
		func(writer http.ResponseWriter, r *http.Request) {
			writer.Write([]byte("hello world"))
		})

	http.ListenAndServe("localhost:8080", nil)
}
```
`http.HandleFunc`是如何实现`http.Handle`一样的功能的呢，下面的源码一目了然。
```
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
```
`http.HandleFunc`转发`DefaultServeMux`的`HandleFunc()`方法。这个方法中其实还是调用了`Handle()`方法。调用`Handle()`方法的前提是将我们传入的函数类型`handler`转成了`HandlerFunc`类型。
```
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	// 空判断
	if handler == nil {
		panic("http: nil handler")
	}

	//  ↓↓↓↓↓↓
	mux.Handle(pattern, HandlerFunc(handler))
}
```
而`HandlerFunc`类型的底层其实还是一个函数类型，并关联了`ServeHTTP()`方法，即实现了`Handler`接口，这样一来，`Handle()`方法便可接收此类型参数了。
```
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```
综上，`http.HandleFunc()`就实现了`http.Handle()`一样的功能。

### 再进一步

#### 几个现成的Handler
为了方便，http包里已经封装好了几个可用的Handler，简单的试一下。
（1）返回404
`http.NotFoundHandler()`这个函数返回的`Handler`就是简单的响应404状态。
```
http.Handle("/nowhere", http.NotFoundHandler())
```
（2）可超时Handler
`http.TimeoutHandler`函数返回的`Handler`存在超时时间，如果在指定的时间没有返回响应就会返回我们设定好消息。  
`http.TimeoutHandler`函数入参有三个：
* http.Handler 类型，即接收请求的`Handler`
* time.Duration 类型，预设的超时时间
* string 类型，超时后返回的消息
```
http.Handle("/timeout", http.TimeoutHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
	}), 1*time.Second, "time out!!!"))
```
（3）执行跳转的Handler
`http.RedirectHandler`函数返回指向Handler就是用来处理状态码3xx的跳转请求。
`http.RedirectHandler`函数入参有两个：
* string 类型，跳转到的Path
* 状态码，一般用3xx
```
http.Handle("/redirect", http.RedirectHandler("b", http.StatusSeeOther))
```
（4）实现文件服务器
通过`http.FileServer(root FileSystem)`函数返回的Handler便可以实现一个文件服务器。其中传入的参数便是文件服务器的根目录。比如`http.FileServer(http.Dir("wwwroot"))`便是指定网站根目录下的“wwwroot”路径作为此文件服务器的根目录。举个例子  
比如现在的目录结构是
> ./  
> main.go   
> wwwroot   
>> test.txt     

现想让wwwroot的路径作为文件服务器的根目录，实现方式很简单
```
package main

import (
    "net/http"
)

func main() {
    http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))
}
```
如果需要实现带路由前缀的文件服务，即比如修改为访问 http://localhost:8080/files 才会路由到文件服务器的话，同样可以使用`http.Handle`函数来设定。如下
```
package main

import (
    "net/http"
)

func main() {
    server := http.Server{
		Addr:    "localhost: 8080",
		Handler: nil, 
	}

	http.Handle("/files/", http.FileServer(http.Dir("wwwroot")))
}
```
但是只是这样写的话，http://localhost:8080/files/test.txt 会返回404状态。原因是当我们访问 http://localhost:8080/files/test.txt 时 handler 会使用设置的根目录路径（wwwroot）拼接上请求的url中的路径（/files/test.txt）即请求的其实是 http://localhost:8080/wwwroot/files/test.txt，这样当然会404。http包中的 `http.StripPrefix`函数便可解决这个问题。它可以帮助我们在使用某个handler时过滤掉url中的一些前缀。  
`http.StripPrefix`函数返回一个handler，入参有两个：
* string类型，即需要过滤的路径前缀
* Handler类型，即需使用的handler，在这里就是http.FileServer返回的handler  
修改后的代码如下
```
package main

import (
    "net/http"
)

func main() {
    server := http.Server{
		Addr:    "localhost: 8080",
		Handler: nil, 
	}

	http.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir("wwwroot"))))
}
```
修改完之后，当访问 http://localhost:8080/files/test.txt 便能正常的显示内容了。