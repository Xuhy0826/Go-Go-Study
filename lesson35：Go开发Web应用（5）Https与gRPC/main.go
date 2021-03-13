package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello"))
	})

	//_ = http.ListenAndServe("localhost:8080", nil)
	//使用https
	err := http.ListenAndServeTLS("localhost:8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}