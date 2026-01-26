package main

import (
	"fmt"
	"net/http"
)

func handleHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Me: I am a Go web server.")
}

func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", handleHelloHandler) // route
	mux.HandleFunc("/about", aboutMeHandler)     // route

	fmt.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
