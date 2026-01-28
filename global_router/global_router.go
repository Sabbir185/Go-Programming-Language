package global_router

import "net/http"

// Middleware
func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
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
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllReq)
}
