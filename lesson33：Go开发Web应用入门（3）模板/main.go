package main

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"text/template"
	"time"
)

func main() {
	//示例1：使用模板
	//useTemplate()

	//示例2：使用多个模板
	//useTemplates()

	//示例3：使用Action
	useAction()
}

//示例1：使用模板
func useTemplate() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("templateFile/info.html")
		_ = t.Execute(w, "hello gopher")
	})

	_ = server.ListenAndServe()
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
		if !strings.HasSuffix(routePath, ".html") {
			routePath = routePath + ".html"
		}
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

	_ = server.ListenAndServe()
}

func loadTmpl() *template.Template {
	tmpls := template.New("tmpls")
	template.Must(tmpls.ParseGlob("templateFile/*.html"))
	return tmpls
}

//示例3：使用Action
var templateCollection *template.Template

func useAction() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	//加载模板集
	templateCollection = loadTmpl()

	http.HandleFunc("/action/", action)

	_ = server.ListenAndServe()
}

//根据不同的请求路径，执行不同的模板
func action(w http.ResponseWriter, r *http.Request) {
	routePath := strings.ToLower(r.URL.Path[8:])
	if !strings.HasSuffix(routePath, ".html") {
		routePath = routePath + ".html"
	}
	//查询模板
	tmpl := templateCollection.Lookup(routePath)
	if tmpl != nil {
		switch routePath {
		case "ifelse.html":
			ifAction(tmpl, w, r)
		case "range.html":
			rangeAction(tmpl, w, r)
		case "with.html":
			withAction(tmpl, w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

//条件Action
func ifAction(t *template.Template, w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	scope := 10
	i := rand.Intn(scope)
	//执行模板
	t.Execute(w, i > scope/2)
}

//迭代Aciton
func rangeAction(t *template.Template, w http.ResponseWriter, r *http.Request) {
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sta", "Sun"}
	t.Execute(w, daysOfWeek)
}

//设置Action
func withAction(t *template.Template, w http.ResponseWriter, r *http.Request) {
	t.Execute(w, "hello")
}
