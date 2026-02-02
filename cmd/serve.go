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

	mux := http.NewServeMux() // router

	initRoutes(mux, manager)

	// globalRouter := middleware.CorsWithPreflight(mux)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
