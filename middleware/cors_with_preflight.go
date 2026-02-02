package middleware

import "net/http"

func CorsWithPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		// handle CORS preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		// other requests
		next.ServeHTTP(w, r)
	})
}
