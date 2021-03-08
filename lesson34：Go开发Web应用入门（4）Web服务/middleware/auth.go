package middleware

import (
	"lesson34/conf"
	"net/http"
)

// AuthMiddleware 授权中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//从配置中读取
		if !conf.App.IsDebug {
			authorization := request.Header.Get("Authorization")
			if authorization != "xuhy" {
				writer.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			next.ServeHTTP(writer, request)
		}
	})
}
