package news

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"lesson34/model"
	"log"
	"net/http"
	"os"
	"strings"
)

// GetNews 查看新闻
func GetNews(w http.ResponseWriter, r *http.Request) {
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
