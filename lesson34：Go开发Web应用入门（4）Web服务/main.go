package main

import (
	"encoding/json"
	"fmt"
	"lesson34/model"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/json", processJson)

	_ = server.ListenAndServe()
}

func processJson(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		serializeJson(w, r)
	} else if r.Method == http.MethodPost {
		deserializeJson(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func serializeJson(w http.ResponseWriter, r *http.Request) {
	post := getModel()
	//创建编码器
	encoder := json.NewEncoder(w)
	//进行编码，成json数据格式
	err := encoder.Encode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

//getModel 返回一个model.Post变量
func getModel() model.Post {
	return model.Post{
		Id:      1,
		Content: "go go go",
		Author: model.Author{
			Id:   2,
			Name: "xuhy",
		},
		Comments: []model.Comment{
			{
				Id:      7,
				Content: "lucky day",
				Author:  "jason",
			},
			{
				Id:      8,
				Content: "what a wonderful life",
				Author:  "jarvis",
			},
		},
	}
}

func deserializeJson(w http.ResponseWriter, r *http.Request) {
	var post = model.Post{}
	//创建解码器
	decoder := json.NewDecoder(r.Body)
	//进行解码，将数据解码到结构上
	err := decoder.Decode(&post)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Printf("%+v", post)
}
