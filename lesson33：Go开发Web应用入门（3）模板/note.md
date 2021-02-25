# Go开发Web应用（3）：模板

Web模板就是预先设计好的html页面，Go的模板引擎会通过模板和传入的数据来创建html页面返回给客户端。Go的标准库`text/template`和`html\template`提供了默认的模板引擎。其实大多时候Go都是作为后端语言来使用，并且大家基本上也都是使用前后端分离的模式搭建应用。但是这也不妨碍我们了解下Go中的模板的用法。

## Go的模板引擎
Go服务器使用模板生成页面进行响应的流程可以大致描述为下图


text/template：通用模板引擎
html/template：Html模板引擎


## 使用模板
使用模板要用到`text/template`包中的`template`结构。首先调用`template`加载相应的模板或者模板集合，随后解析模板并根据传入的数据生成html写入响应。主要涉及到的两个方法是加载模板和执行模板（生成响应）。  
`template`使用`ParseFiles`方法读取和解析模板（集合），方法返回指向解析后的`template.Template`的指针与一个可为空的错误。之后使用`Execute`方法执行模板生成html。  
【示例1】使用模板
```go
package main

import (
	"net/http"
	"text/template"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		//**********************************************************
		t, _ := template.ParseFiles("templateFile/info.html")
		t.Execute(w, "hello gopher")
		//**********************************************************
	})

	server.ListenAndServe()
}
```
info.html模板文件如下
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Info</title>
</head>

<body>
    {{.}}
</body>

</html>
```



### action
#### 条件类
```
{{if arg}}
...
{{else}}
...
{{end}}
```

#### 迭代遍历
可以遍历array，slice，map和channel等。结构如下
```
{{range array}}
Dot is set to the element{{.}}
{{end}}
```
“.”就表示每次循环的元素。

#### 设置类


#### 包含类
```
{{template "name"}}
或
{{template "name" arg}}
```

#### 定义类
define action