package main

import (
	"lesson34/middleware"
	"lesson34/router"
	"net/http"
)

func main() {
	r := router.New(router.Routes...)

	//调用顺序： logging -> auth -> next
	r.Use(middleware.LoggingMiddleware, middleware.AuthMiddleware)

	_ = http.ListenAndServe("localhost:8080", r)
}
