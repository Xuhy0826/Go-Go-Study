package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"lesson34/model"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)

	r.HandleFunc("/post", getPostHandler).
		Methods(http.MethodGet). //限制访问方法：POST
		Schemes("http").         //设置scheme为http
		Name("getpost")          //命名路由

	r.Path("/post").
		Methods(http.MethodPost).                   //限制访问方法：POST
		HandlerFunc(postPostHandler).               //设置处理方法
		Schemes("http").                            //设置scheme为http
		Headers("Content-Type", "application/json") //设置请求头

	r.HandleFunc("/news/{title:[a-z]+}", newList). //路由参数：支持正则
							Methods(http.MethodGet)

	_ = http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome!"))
}

// getPostHandler 获取一个帖子的内容
func getPostHandler(w http.ResponseWriter, r *http.Request) {
	serializeJson(w, r)
}

// postPostHandler 上传一个帖子的内容
func postPostHandler(w http.ResponseWriter, r *http.Request) {
	deserializeJson(w, r)
}

func newList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //获取参数，map类型
	var newsList []model.News

	allnews := getAllNews() //获取所有 news 集合
	for _, news := range allnews {
		if strings.Contains(news.Title, vars["title"]) {
			newsList = append(newsList, news)
		}
	}
	//json 序列化后返回
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&newsList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

// 读取所有新闻列表
func getAllNews() []model.News {
	filePtr, err := os.Open("data/new-list.json")
	if err != nil {
		log.Fatalf("Open file failed [Err:%s]\n", err.Error())
	}
	defer filePtr.Close()
	var newsList []model.News

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&newsList)
	if err != nil {
		log.Fatal("Decoder failed", err.Error())
		return nil
	} else {
		return newsList
	}
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
