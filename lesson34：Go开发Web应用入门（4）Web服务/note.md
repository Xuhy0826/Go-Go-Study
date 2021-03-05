# Go开发Web应用（4）：服务

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

下面在Handler中读取请求中传来的json数据

