# WebDemo(1)

`http.ListenAndServe("localhost:8080", nil)`实际上是创建一个`http.Server`并调用其`ListenAndServe()`函数。
> `http.Server` 结构定义
```
type Server struct {
    //网络地址，为空则默认为“*:80"
	Addr string

    //默认为http.DefaultServeMux
	Handler Handler
}
```
> `ListenAndServe()`函数声明
```
func (srv *Server) ListenAndServe() error {
	if srv.shuttingDown() {
		return ErrServerClosed
	}
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}
```
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

### DefaultServeMux
`DefaultServeMux`是一个默认的`ServeMux`，同时也实现了 `Handler` 接口
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
`Handler`接口就是用来处理请求，针对不同的请求地址应该制定不同的`Handler`进行处理，`DefaultServeMux`来作为前置的`Handler`来分发请求，相当于一个路由器。  
![DefaultServeMux路由请求到各个Handler](https://github.com/Xuhy0826/Golang-Study/blob/master/resource/httpHandler.png)