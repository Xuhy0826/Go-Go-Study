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
		Items []gResult `json:"items"`
		Kind  string    `json:"kind"`
	}

	gResult struct {
		Kind             string `json:"kind"`
		Title            string `json:"title"`
		HTMLTitle        string `json:"htmlTitle"`
		Link             string `json:"link"`
		DisplayLink      string `json:"displayLink"`
		Snippet          string `json:"snippet"`
		HTMLSnippet      string `json:"htmlSnippet"`
		FormattedURL     string `json:"formattedUrl"`
		HTMLFormattedURL string `json:"htmlFormattedUrl"`
		Mime             string `json:"mime"`
		FileFormat       string `json:"fileFormat"`
	}
)

func jsonTest() {
	uri := "https://www.googleapis.com/customsearch/v1/siterestrict?key=AIzaSyCIivhVfq-5L9yT8RQ9J8olrRV67lE_Ta8&cx=017576662512468239146:omuauf_lfve&q=golang"

	//向api发起搜索，得到响应
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	var gr gResponse

	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//本例的核心：将响应的json文档解码到声明的结构体中
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Printf("%+v\n", gr)
}

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Address struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"address"`
}

func jsonDeserilizeTest() {
	var JSON = `{
		"name": "Gopher",
		"title": "programmer",
		"address": {
			"home": "415.333.3333",
			"cell": "415.555.5555"
		}
	}
	`
	var c Contact
	//进行json反序列化
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("Error", err)
		return
	}
	fmt.Printf("%+v", c)
}

func jsonDeserilizeTest2() {
	var JSON = `{
		"name": "Gopher",
		"title": "programmer",
		"address": {
			"home": "415.333.3333",
			"cell": "415.555.5555"
		}
	}
	`
	var c map[string]interface{}
	//进行json反序列化
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("Error", err)
		return
	}

	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Address")
	fmt.Println("H:", c["address"].(map[string]interface{})["home"])
	fmt.Println("C:", c["address"].(map[string]interface{})["cell"])
}

func jsonSerilizeTest() {
	contact := Contact{
		Name:  "Gopher",
		Title: "programmer",
	}
	contact.Address.Home = "415.333.3333"
	contact.Address.Cell = "415.555.5555"

	data, err := json.MarshalIndent(contact, "", " ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))
}
