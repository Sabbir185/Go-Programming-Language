package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	// common middleware
	manager := middleware.NewManager()
	manager.Use(middleware.CorsWithPreflight, middleware.Logger, middleware.GeneralPrint)

	// HTTP Multiplexer --> router
	mux := http.NewServeMux()

	initRoutes(mux, manager)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
