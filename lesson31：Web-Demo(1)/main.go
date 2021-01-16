package main

import "net/http"

func main() {
	//startDefaultServer()

	//startMyServer()

	multiHandlerServer()
}

func startDefaultServer() {
	//第一个参数：路由地址
	//第二个参数：处理请求的函数
	http.HandleFunc("/",
		func(writer http.ResponseWriter, r *http.Request) {
			writer.Write([]byte("hello world"))
		})

	//开启http请求监听
	//参数1 ：地址，默认为 *:80
	//参数2 ：handler，默认是 DefaultServeMux，是一个multiplexer（看成路由器）
	http.ListenAndServe("localhost:8080", nil)
}

func startMyServer() {
	//自定义server
	server := http.Server{
		Addr:    "localhost: 8080",
		Handler: myHandler{}, //不指定，默认为 DefaultServeMux
	}

	server.ListenAndServe()
}

// 自定义 Handler
type myHandler struct{}

// myHandler 实现 Handler 接口
// 接收请求，返回"Hello gopher"字符串
func (h myHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello gopher"))
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

type aHandler struct{}

func (ah aHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello gopher from aHandler"))
}

type bHandler struct{}

func (bh bHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello gopher from bHandler"))
}
