package main

import "net/http"

func main() {

}

func startDefaultServer() {
	//第一个参数：路由地址
	//第二个参数：处理请求的函数
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("hello world"))
	})

	//开启http请求监听
	//参数1 ：地址，默认为 *:80
	//参数2 ：handler，默认是 DefaultServeMux，是一个multiplexer（看成路由器）
	http.ListenAndServe("localhost:8080", nil)
}

func startMyServer() {

	server := http.Server{
		Addr:    "localhost: 8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
