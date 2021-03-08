package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware 日志中间件
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if next == nil {
			next = http.DefaultServeMux
		}
		log.Printf("request from %v", request.RemoteAddr)
		next.ServeHTTP(writer, request)
	})
}
