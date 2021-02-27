# Go开发Web应用（3）：模板

Web模板就是预先设计好的html页面，Go的模板引擎会通过模板和传入的数据来创建html页面返回给客户端。Go的标准库`text/template`和`html\template`提供了默认的模板引擎。其实大多时候Go都是作为后端语言来使用，并且大家基本上也都是使用前后端分离的模式搭建应用。但是这也不妨碍我们了解下Go中的模板的用法。

## Go的模板引擎
GO的模板就是一个文本文件（常为html文件），其中嵌入一些action。如果Go服务器使用模板生成页面，那么流程简单描述：Handler调用模板引擎，将一个或多个模板文件传入给模板引擎，在传入模板需要的动态数据；模板在收到这些数据后会生成相应的html文件，并将这些html文件写入到`ResponseWriter`中最终返回给客户端。可以大致描述为下图。
![Go服务器使用模板生成页面](https://github.com/Xuhy0826/Golang-Study/blob/master/resource/goTemplateFlow.jpg)  

### 使用模板
使用模板要用到`text/template`包中的`template`结构。首先调用`template`加载相应的模板或者模板集合，随后解析模板并根据传入的数据生成html写入响应。主要涉及到的两个方法是加载模板和执行模板（生成响应）。  
`template`使用`ParseFiles`方法读取和解析模板（集合），方法返回指向解析后的`template.Template`的指针与一个可为空的错误。接着使用`Execute`方法执行模板将数据应用到模板中生成html。  
比如下面的info.html就是一个模板文件。和真正的html文件的区别就是其中有使用`{{ }}`包住的内容。`{{ }}`包住的内容区域就是前面所说的Action，模板引擎在执行模板是会使用一个值去替换这个Action本身。如果你用过使用过其他语言进行过动态网页的开发应该很熟悉这种模式，对于我来说联想到的就是Razor页面。
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
这个模板的Action只有一个"."，表示模板引擎在执行模板时，使用一个值去替换这个Action本身。有了模板文件，加载并执行也并不复杂。如下所示
##### 【示例1】使用模板

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
		t, _ := template.ParseFiles("templateFile/info.html") //使用了相对路径
		_ = t.Execute(w, "hello gopher")    //使用字符串"hello gopher"来替换
		//**********************************************************
	})

	_ = server.ListenAndServe()
}
```
`template.ParseFiles`是包提供的独立函数，其背后是调用了`Template`结构的`ParseFiles`方法。
```go
t, _ := template.ParseFiles("tmpl.html")
```
这和下面的写法是等价的
```go
t := template.New("tmpl.html")
t, _ := t.ParseFiles("tmpl.html")
```
`ParseFiles`方法可以接收多个参数，其函数签名如下
```go
func parseFiles(t *Template, readFile func(string) (string, []byte, error), filenames ...string) (*Template, error) {
	......
}
```
无论接收几个文件名做参数，都只返回一个模板。并且这个模板的名称就是传入的第一个文件的文件名（带后缀）。但是如果传入的是多个模板文件，那么会将这些模板文件放在一个集合里保存。  
除了`ParseFiles`函数，另外还有一个`ParseGlob`函数也可对模板进行语法分析。`ParseGlob`函数传入的参数是带通配符的文件名。比如`t.ParseGlob("templateFile/*.html)`就是将templateFile文件夹内所有后缀为.html的文件一起传入进行分析。并且可以通过Lookup方法根据模板名找出需要的模板进行执行。
```go
templateCollection := template.ParseGlob("templateFile/*.html")
tmpl := templateCollection.Lookup("a.html")
tmpl.Execute(w, nil)
```
另外，`template.Must`函数提供了一种处理错误的机制。`template.Must`可以包裹一个待执行的函数，待执行的函数返回的是一个指向模板的指针和一个可为空的错误。如果这个错误不是nil，那么`template.Must`函数将会引发`panic`。如果没有错误，返回模板指针。
```go
t := template.Must(template.ParseGlob("templateFile/*.html))
```

### action
在模板文件中使用双大括号包裹的内容就称为Action，单纯的点“.”就是一个简单的Action。当然除了“.”还有其他的Action。

#### 条件Action
条件Action根据参数值来进行分支操作。
```
{{if arg}}
...
{{else}}
...
{{end}}
```
当然，可以没有“else”分支。一个简单的示例即可说明用法。
##### 【示例2】条件Action的使用
```go
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/action", action)

	_ = server.ListenAndServe()
}

func action(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templateFile/ifelse.html")
	//产生随机数
	rand.Seed(time.Now().Unix())
	scope := 10
	i := rand.Intn(scope)
	//执行模板
	t.Execute(w, i > scope/2)
}
```
模板文件如下，只包含一层条件Action。由于传入的参数是由随机数控制，所以不停访问会随机返回两种相应。
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>IF-ELSE</title>
</head>

<body>
    {{if .}}
    Lorem ipsum dolor sit amet consectetur adipisicing elit. Facere sit aut quos, natus iure alias dolore quam! Eum hic,
    fuga quasi eos impedit distinctio nam iusto commodi nemo odit libero.
    {{else}}
    襟三江而带五湖，控蛮荆而引瓯越。物华天宝，龙光射牛斗之墟；人杰地灵，徐孺下陈蕃之榻。雄州雾列，俊采星驰。台隍枕夷夏之交，宾主尽东南之美。
    {{end}}
</body>

</html>
```

#### 迭代Action
可以遍历array，slice，map和channel等。结构如下，其中迭代内部的点（.）会被赋予被迭代的元素。
```
{{range array}}
Dot is set to the element{{.}}
{{else}}
Nothing to show
{{end}}
```
如果加上`{{else}}`段，那么传入的集合为`nil`时，则会进入`{{else}}`段。
##### 【示例3】迭代Action的使用
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Range</title>
</head>

<body>
    <ul>
    {{range .}}
    <li>{{.}}</li>
    {{else}}
    <li>nothing</li>
    {{end}}
</ul>
</body>

</html>
```
Handler中向模板中传入一个string数组进行迭代
```go
func rangeAction(t *template.Template, w http.ResponseWriter, r *http.Request) {
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sta", "Sun"}
	t.Execute(w, daysOfWeek)
}
```

#### 设置Action
设置Action允许用户在指定范围内为点（.）设置值。就是在{{with %你想设置的值%}}与{{end}}之间，点（.）不再是传入的值，而是你设置的值。
```
{{with arg}}
 Dot is set to arg
{{end}}
```
用个例子说明。
##### 【示例4】设置Action的用法
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>With </title>
</head>

<body>
<div>点表示： {{.}}</div>
<div>
    {{with "gopher"}}
    点表示： {{.}}
    {{end}}
</div>
</body>

</html>
```
```go
//设置Action
func withAction(t *template.Template, w http.ResponseWriter, r *http.Request) {
	t.Execute(w, "hello")
}
```
#### 包含Action
```
{{template "name"}}
或
{{template "name" arg}}
```

#### 定义类
define action