package post

import (
	"encoding/json"
	"fmt"
	"lesson34/model"
	"net/http"
)

// GetPost 获取一个帖子的内容
func GetPost(w http.ResponseWriter, r *http.Request) {
	serializeJson(w, r)
}

// PostPost 上传一个帖子的内容
func PostPost(w http.ResponseWriter, r *http.Request) {
	deserializeJson(w, r)
}

// 返回一个 帖子 的json数据
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

// 获取上传的json数据，反序列成对象
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
