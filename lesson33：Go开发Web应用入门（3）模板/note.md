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
包含Action可以让模板实现嵌套。写法是下面这样
```
{{template "name"}}
或
{{template "name" arg}}
```
其中`name`就是要包含的模板的名称。下面这个模板"t1.html"中嵌套了"t2.html"。前面说过模板文件的名称会被用作为模板的名。
```html
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>t1</title>
</head>
	<body>
		<div>This is t1.html before</div>
		<div>This is the value of the dot in t1.html - {{.}}</div>
		<hr/>
		{{template "t2.html"}}
		<hr/>
		<div>This is t1.html after</div>
	</body>
</html>
```

而"t2.html"模板如下
```html
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>t2</title>
</head>
<body>
	<div style="background-color: dodgerblue">
		This is t2.html
		This is the value of the dot in t2.html - {{.}}
	</div>
</body>
</html>
```
访问 `http://localhost:8080/action/t1` 能看到下面的界面。  
![包含Action](https://github.com/Xuhy0826/Golang-Study/blob/master/resource/t1html.jpg)

需要注意的是，加载嵌套的模板需要对所有涉及到的模板进行分析，就是在调用`ParseFiles`函数或者`ParseGlob`函数时要将涉及到的模板都传入。上面模板t1.html中的点(.)被传入的字符串所替换，但是t2.html中的点(.)却没有。如果将t1.html中的Action改写成`{{template "name" arg}}`这样便可以将变量传递到嵌套的模板中，如下
```html
{{template "t2.html" .}}
```
现在访问 `http://localhost:8080/action/t1` 的界面如下。
![包含Action](https://github.com/Xuhy0826/Golang-Study/blob/master/resource/t1v2html.jpg)

### 参数、变量和管道
#### 参数和变量
前面一直在模板中使用的点(.)就是一个参数，表示的是Handler向模板传递的数据。除了点(.)，参数还可以是bool、int、string等字面量，也可以是结构或者方法，但是方法只能有一个返回值或者一个返回值加一个可为空的错误。  
用户还可以在模板中定义变量，使用`$`符号开头，如下
```html
{{range $key, $value := .}}
The key is {{$key}} and the value is {{$value}}
{{end}}
```
这样，从Handler传给模板一个map时，便可以进行遍历了。

#### 管道
模板中的管道是多个有序的串联起来的参数、函数或者方法。工作方式和Linux中的管道有点类似。
```
{{p1 | p2 | p3}
```
在管道中，p1的输出作为p2的输入，依次下去。举个简单的例子
```
{{12.3456 | printf "%.2f"}}
```
上面的管道，浮点数字面量作为参数输入到模板的内置函数`printf`中，并使用指定的格式符，最终输出12.35.

### 函数
Go的函数可以作为参数输入模板，并且Go模板引擎也内置了一些函数，这些函数都有限制：只能有一个返回值或者一个返回值加一个可为空的错误。  
用户创建自定义的模板函数的步骤：  
（1）创建一个`FuncMap`的映射，并将映射的键设置为函数的名字，值设置成实际函数
（2）将`FuncMap`与模板进行绑定
【示例5】自定义模板函数
```go
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}

//模板函数的使用
func process(tw http.ResponseWriter, r *http.Request) {
	//step1: 创建FuncMap映射
	funcMap := template.FuncMap{
		"fdate": formatDate,
	}
	//step2: 将FuncMap映射与模板关联
	t := template.New("tmpl.html").Funcs(funcMap)

	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

//自定义模板函数：格式化日期
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
```
模板文件如下
```html
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Process</title>
</head>
<body>
	<div>The date time is {{ . | fdate }}</div>
	<!--或者-->
	<div>The date time is {{ fdate . }}</div>
</body>
</html>
```
但是，综合来说，管道还是会比函数要更强大和灵活，并更加易读。

### 布局Layout
布局页（layout）在很多其他的Web框架中也经常见到，比如ASP.NET(Core)中也有layout的使用。使用`{{template "name" .}}`可以在模板中实现嵌套，如果使用这样的方法来实现布局页，那么每个页面都需要有一个自己的布局页文件，那么意义就不是很大了。我们需要的是一个公共的布局页。这里，Go还提供了定义Action来帮助实现布局页。  
【示例6】使用布局页
```html
{{define "layout"}}

<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Go Layout</title>
</head>
<body>
    <div>This is Layout - begin</div>
    <div>this is value - {{.}}</div>
    {{template "content"}}
    <div>This is Layout - end</div>
</body>
</html>

{{end}}
```
```html
{{define "content"}}

<h1 style="color: red">Hello world</h1>

{{end}}
```
```html
{{define "content"}}

<h1 style="color: blue">Hello world</h1>

{{end}}
```
下面是后台代码
```go
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/layout", layout)
	_ = server.ListenAndServe()
}

func layout(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = t.ParseFiles("layout.html", "redHello.html")
	} else {
		t, _ = t.ParseFiles("layout.html", "blueHello.html")
	}
	_ = t.ExecuteTemplate(w, "layout", "")
}
```
和之前不一样的是，这里使用的执行模板的方法是`ExecuteTemplate`，并把待执行的模板名传入。这样的写法便实现了布局页。