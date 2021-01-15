package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes 为网络服务设置路由
func Routes() {
	http.HandleFunc("/sendjson", func(rw http.ResponseWriter, r *http.Request) {
		u := struct {
			Name  string
			Email string
		}{
			Name:  "xuhy",
			Email: "xuhy@github.com",
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		json.NewEncoder(rw).Encode(&u)
	})
}
