package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	//根据api返回的相应的json的结构定义结构体
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}

	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}
)

func jsonTest() {
	uri := "https://customsearch.googleapis.com/customsearch/v1?cx=823e7135ee68a70b7&q=GO-GO-Study&key=AIzaSyCIivhVfq-5L9yT8RQ9J8olrRV67lE_Ta8"

	//向api发起搜索，得到响应
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	var gr gResponse

	//本例的核心：将响应的json文档解码到声明的结构体中
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}
