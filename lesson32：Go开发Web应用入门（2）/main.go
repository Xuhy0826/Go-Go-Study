package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintln(w, r.Form)
	})

	http.ListenAndServe("localhost:8080", nil)
}
