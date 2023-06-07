package middleware

import (
	"fmt"
	"net/http"
)

type Func func(handler http.Handler) http.Handler

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Path: %s, Method: %s\n", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}
