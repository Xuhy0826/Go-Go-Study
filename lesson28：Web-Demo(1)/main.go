package main

import "net/http"

func main() {
	//第一个参数：路由地址
	//第二个参数：处理请求的函数
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("hello world"))
	})

	//开启http请求监听
	http.ListenAndServe("localhost:8080", nil) //DefaultServerMux
}
