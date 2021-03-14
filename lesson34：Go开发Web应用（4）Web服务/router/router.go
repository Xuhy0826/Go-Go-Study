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
