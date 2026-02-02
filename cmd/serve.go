package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	// HTTP Multiplexer --> router
	mux := http.NewServeMux()

	// common middleware
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	wrapMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", wrapMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
