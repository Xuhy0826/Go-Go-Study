# Go开发Web应用（4）：Web服务

前言：通过前面的学习，了解了使用的Go来进行Web编程的一些基本知识。如何开发一个有一定规范性的的Web服务仍有很多地方需要改进，比如开发一个REST API，如何处理不同的http方法，数据传输，权限验证和配置等等如何处理呢。

> 本文示例代码已上传至GitHub：[传送门](https://github.com/Xuhy0826/Go-Go-Study/tree/master/lesson34%EF%BC%9AGo%E5%BC%80%E5%8F%91Web%E5%BA%94%E7%94%A8%E5%85%A5%E9%97%A8%EF%BC%884%EF%BC%89Web%E6%9C%8D%E5%8A%A1)

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

接下来简单介绍下基于gorilla/mux来配置路由的方法，安装直接使用`go get`命令即可，切到自己的工作目录下执行

```bash
go get -u github.com/gorilla/mux
```

首先`gorilla/mux`的最基本用法如下

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

以上的代码中，所有对路由的配置工作都堆在了main函数中，既不优雅也不利于以后的维护。而比较实用的做法是将各个功能组的路由分到不同的go文件中，在使用同一的配置入口进行注册。就以上面所用到的代码为例，将所有的``Handler`都分类单独创建go文件保存，并创建一个统一的配置入口来进行路由的配置。

创建一个`/router`的路径，上面的示例代码涉及到帖子（post）和新闻（news）两个分类的服务，再创建`/router/post`与`/router/news`两个路径来分别存放相应的go文件。单独抽出来的文件比如以news.go文件为例，代码如下

1. /router/news/news.go

```go
package news

import ......

// GetNews 查看新闻
func GetNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //获取参数，map类型
	var newsList []model.News

	allnews := getAllNews() //获取所有 news 集合
	......
}

// 读取所有新闻列表
func getAllNews() []model.News {
	......
	
}

```

这样一来，所有的`Handler`结构就清晰了很多。接下来为创建一个配置的统一入口，为了方便配置，创建一个`Route`类型来设置配置信息。

- /router/router.go

```go
package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route 路由的配置信息
type Route struct {
	Name        string
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

// 创建Router
func New(routeCollection ...Route) *mux.Router {
	router := mux.NewRouter()
	//router.StrictSlash(true)

	for _, route := range routeCollection {
		router.Path(route.Path).
			Name(route.Name).
			Methods(route.Method).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}
```

最后，将项目用到的所有路由配置也单独存放于一个go文件中

```go
package router

import (
	"lesson34/router/index"
	"lesson34/router/news"
	"lesson34/router/post"
)

var Routes = []Route{
	{
		Name:        "index",
		Path:        "/",
		Method:      "GET",
		HandlerFunc: index.Index,
	},
	{
		Name:        "get_post",
		Path:        "/post",
		Method:      "GET",
		HandlerFunc: post.GetPost,
	},
	{
		Name:        "post_post",
		Path:        "/post",
		Method:      "POST",
		HandlerFunc: post.PostPost,
	},
	{
		Name:        "news",
		Path:        "/news/{title}",
		Method:      "GET",
		HandlerFunc: news.GetNews,
	},
}
```

这样一来，在main函数中，我们 只需调用路由的统一配置方法即可，整个项目的结构变得清晰也利于后期的维护管理。

```go
package main

import (
	"lesson34/router"
	"net/http"
)

func main() {
	r := router.New(router.Routes...)
	_ = http.ListenAndServe("localhost:8080", r)
}
```

重构好的文件结构会是下面这样，结构清晰。

```bash
│  go.mod
│  go.sum
│  main.go
├─data
│      new-list.json
├─model
│      Author.go
│      news.go
│      post.go
└─router
    │  routecfg.go
    │  router.go
    ├─index
    │      index.go
    ├─news
    │      news.go
    └─post
           post.go
```

## 中间件

中间件的概念应该都不会陌生，顾名思义，中间件就是在请求处理管道中增加的处理环节。用处很多，比如实现记录日志，身份验证，超时处理等等。在Go中实现一个中间件是很简单的，因为Go处理请求的单元是一个个`Handler`，那么中间件其实也是一个`Handler`。只需要将下一个调用的`Handler`作为参数传入，那么就可以实现了。

比如创建一个记录日志的`Handler`

```go
package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware 日志中间件
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if next == nil {
			next = http.DefaultServeMux
		}
		log.Printf("request from %v", request.RemoteAddr)
		next.ServeHTTP(writer, request)
	})
}
```

main函数在启动监听时，只需将中间件加上即可。

```go
package main

import (
	"lesson34/middleware"
	"lesson34/router"
	"net/http"
)

func main() {
	r := router.New(router.Routes...)
	_ = http.ListenAndServe("localhost:8080", middleware.LoggingMiddleware(r))
}
```

实现多个中间件只需按照同样的思路进行串联即可。而`gorilla/mux`组件中也包含了中间件功能的封装，使用`r.Use()`函数便可以为全局设置中间件。

```go
func main() {
	r := router.New(router.Routes...)

	//使用中间件
	r.Use(middleware.LoggingMiddleware)

	_ = http.ListenAndServe("localhost:8080", r)
}
```

如果有多个中间件的情况，由于`r.Use()`函数的入参是可接受多个参数的，只需按照调用顺序传入各个中间件即可。

比如我们有另外一个中间件`AuthMiddleware`。

```go
package middleware

import (
	"net/http"
)

// AuthMiddleware 授权中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		authorization := request.Header.Get("Authorization")
		if authorization != "xuhy" {
			writer.WriteHeader(http.StatusUnauthorized)
		}else {
			next.ServeHTTP(writer, request)
		}
	})
}
```

在main函数中，使用use函数将两个中间件依次传入

```go
package main

import (
	"lesson34/middleware"
	"lesson34/router"
	"net/http"
)

func main() {
	r := router.New(router.Routes...)

	//调用顺序： logging -> auth -> next
	r.Use(middleware.LoggingMiddleware, middleware.AuthMiddleware)

	_ = http.ListenAndServe("localhost:8080", r)
}
```



