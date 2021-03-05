package router

import "net/http"

// Route 用来定义一个路由
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter(routes []Route) *http.Handler {
	for _, route := range routes {

	}
}

type Router struct {
	handler http.Handler

	namedRoutes map[string]*Route
}

type matcher interface {
	Match(*http.Request, *RouteMatch) bool
}

type RouteMatch struct {
	Route   *Route
	Handler http.Handler
	Vars    map[string]string
}
