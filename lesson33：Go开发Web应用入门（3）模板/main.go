package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	//示例1：使用模板
	//useTemplate()

	//示例2：使用多个模板
	useTemplates()
}

//示例1：使用模板
func useTemplate() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("templateFile/info.html")
		t.Execute(w, "hello gopher")
	})

	server.ListenAndServe()
}

//示例2：使用多个模板
func useTemplates() {
	//加载所有模板
	templateCollection := loadTmpl()

	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//根据请求路径取得相应的模板名
		routePath := r.URL.Path[1:]
		tmpl := templateCollection.Lookup(routePath)
		if tmpl != nil {
			err := tmpl.Execute(w, nil)
			if err != nil {
				log.Fatalln(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/imgs/", http.FileServer(http.Dir("wwwroot")))

	server.ListenAndServe()
}

func loadTmpl() *template.Template {
	tmpls := template.New("tmpls")
	template.Must(tmpls.ParseGlob("templateFile/*.html"))
	return tmpls
}
