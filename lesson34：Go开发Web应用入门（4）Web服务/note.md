# Go开发Web应用（4）：Web服务

## 处理Json数据

在与REST Web服务进行交互时，使用最频繁的数据格式应该就是JSON了。前面我们也了解了 [Go如何（反）序列化对象](https://my.oschina.net/xuhy0826/blog/4956986)，通过使用内置的json包便可以实现。除此之外，还可以使用encoding包来实现对json数据的处理，而且encoding包更方便在Web编程中使用，因为它是根据编码器或者解码器来处理**流式数据**。

### 读取Json数据

使用encoding包来读取json数据的大致流程可以分为3步

1. 创建用户存储json数据的结构
2. 创建出用于解码的json数据的解码器
3. 遍历json数据将其解码到结构中

比如下面定义的结构

```go
package model

// Post 表示论坛中的帖子
type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

// Comment 表示帖子的评论
type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Author 表示帖子或者评论的作者
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

```

由于`json.NewDecoder()`函数接收的参数是`io.Reader`接口类型，而`request.Body`恰好满足`io.Reader`接口。下面示例展示了在Handler中读取请求中传来的json数据。

```go
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
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
	}
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

```

### 创建json

相对应的，使用`encoding/json`来创建json数据返回也是类似的流程，需要通过`json.NewEncoder()`函数来得到编码器，入参需要`io.Writer`接口，而恰好`http.ResponseWriter`满足接口。再将需要序列化的结构作为参数传入编码器的进行编码从而得到json数据。

以下示例展示了这个过程。

```go
func processJson(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
        // 创建json数据返回
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
```

## 路由

搭建REST API Service需要为不同的请求路径和请求方式设置不同的处理流程，也就是指定相应的Handler。比如之前开发ASP.NET Core WebApi 使用MVC的模式，框架基本已经帮我们封装路由功能，只需简单配置与标注特性就可以自定义路由规则。在Go中使用`DefaultServerMux`也可以实现最基本的路由功能，当然有很多开源的轮子用起来能更加方便，比如[Gorilla/Mux](https://github.com/gorilla/mux,"Gorilla/Mux")或者 [httprouter](https://github.com/julienschmidt/httprouter)。

接下来简单介绍下基于gorilla/mux来配置路由的方法，首先gorilla/mux的最基本用法如下

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"lesson34/model"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/post", getPostHandler)

	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome!"))
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	serializeJson(w, r)
}
```

初看起来和之前使用标准库中的写法差不多，只是把`DefaultServerMux`替换成了`mux.Router`。接下来看下如何进行更多细节的配置。

````go
r.HandleFunc("/post", getPostHandler).
	Methods(http.MethodGet). //限制访问方法：POST
	Schemes("http").         //设置scheme为http
	Name("getpost")          //命名路由

r.Path("/post").
	Methods(http.MethodPost).                   //限制访问方法：POST
	HandlerFunc(postPostHandler).               //设置处理方法
	Schemes("http").                            //设置scheme为http
	Headers("Content-Type", "application/json") //设置请求头
````

除此之外，我们还可以进行路由参数的配置，而且支持正则的匹配

```go
r.HandleFunc("/news/{title:[a-z]+}", newList). //路由参数：支持正则
	Methods(http.MethodGet)
```

在`Handler`中获取路由参数的方法如下

```go
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
```

以上简单介绍了`gorilla/mux`中路由的使用，更多的详细说明还是参阅官方文档靠谱。

以上的代码中，所有对路由的配置工作都堆在了main函数中，既不优雅也不利于以后的维护。而比较实用的做法是将各个功能组的路由分到不同的go文件中，在使用同一的配置入口进行注册。就以上面所用到的代码为例，

