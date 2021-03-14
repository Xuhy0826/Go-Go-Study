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
	//useAction()

	//示例5：自定义模板函数
	//useFunc()

	//示例6：使用布局页
	//useLayout()
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
		case "t1.html":
			NestingAction(tmpl, w, r)
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
	_ = t.Execute(w, i > scope/2)
}

//迭代Aciton
func rangeAction(t *template.Template, w http.ResponseWriter, r *http.Request) {
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sta", "Sun"}
	t.Execute(w, daysOfWeek)
}

//【示例4】设置Action
func withAction(t *template.Template, w http.ResponseWriter, r *http.Request) {
	_ = t.Execute(w, "hello")
}

//包含Action
func NestingAction(t *template.Template, w http.ResponseWriter, r *http.Request) {
	_ = t.Execute(w, "Hello gopher")
}

//示例5：模板函数的使用
func useFunc() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}

//模板函数的使用
func process(w http.ResponseWriter, r *http.Request) {
	//step1: 创建FuncMap映射
	funcMap := template.FuncMap{
		"fdate": formatDate,
	}
	//step2: 将FuncMap映射与模板关联
	t := template.New("tmpl.html").Funcs(funcMap)

	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

//自定义模板函数
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

//示例6：layout的使用
func useLayout() {
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
