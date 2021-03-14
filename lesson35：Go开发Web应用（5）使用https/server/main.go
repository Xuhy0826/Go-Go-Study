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
