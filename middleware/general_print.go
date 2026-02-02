package middleware

import (
	"fmt"
	"net/http"
)

func GeneralPrint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("General Print Middleware: Received request for", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
