# Go开发Web应用（2）：模板

类似Razor页面

text/template：通用模板引擎
html/template：Html模板引擎

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
		t, _ := template.ParseFiles("templateFile/info.html")
		t.Execute(w, "hello gopher")
	})

	server.ListenAndServe()
}
```
模板文件如下
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