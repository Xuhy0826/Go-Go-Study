package main

import (
	"fmt"
	"lesson34/conf"
	"lesson34/middleware"
	"lesson34/router"
	"net/http"
)

func main() {
	//加载配置
	conf.Init("app.toml")

	r := router.New(router.Routes...)
	//调用顺序： logging -> auth -> next
	r.Use(middleware.LoggingMiddleware, middleware.AuthMiddleware)

	addr := fmt.Sprintf("%s:%s", conf.App.SeHost, conf.App.SePort)
	_ = http.ListenAndServe(addr, r)
}
