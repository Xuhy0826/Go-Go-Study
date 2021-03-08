package middleware

import (
	"net/http"
)

// AuthMiddleware 授权中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		authorization := request.Header.Get("Authorization")
		if authorization != "xuhy" {
			writer.WriteHeader(http.StatusUnauthorized)
		} else {
			next.ServeHTTP(writer, request)
		}
	})
}
