package main

import "net/http"

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello"))

	})

	_ = http.ListenAndServe("localhost:8080", nil)
}