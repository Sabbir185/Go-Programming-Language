package middleware

import (
	"fmt"
	"net/http"
)

func GeneralPrint(next http.Handler) http.Handler {
	return http.HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Route executed")
			next.ServeHTTP(w, r)
		}
	)
}
