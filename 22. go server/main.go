package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Me: I am a Go web server.")
}

type Products struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"imgUrl"`
}

var productlist []Products

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Request, not found!", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productlist)
}

func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", handleHelloHandler) // route
	mux.HandleFunc("/about", aboutMeHandler)     // route
	mux.HandleFunc("/products", getProductHandler)

	fmt.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	p1 := Products{
		ID:          1,
		Title:       "Product 1",
		Price:       19.99,
		Description: "This is the first product.",
		ImgUrl:      "http://example.com/product1.jpg",
	}
	p2 := Products{
		ID:          2,
		Title:       "Product 2",
		Price:       29.99,
		Description: "This is the second product.",
		ImgUrl:      "http://example.com/product2.jpg",
	}
	p3 := Products{
		ID:          3,
		Title:       "Product 3",
		Price:       39.99,
		Description: "This is the third product.",
		ImgUrl:      "http://example.com/product3.jpg",
	}
	p4 := Products{
		ID:          4,
		Title:       "Product 4",
		Price:       49.99,
		Description: "This is the fourth product.",
		ImgUrl:      "http://example.com/product4.jpg",
	}
	p5 := Products{
		ID:          5,
		Title:       "Product 5",
		Price:       59.99,
		Description: "This is the fifth product.",
		ImgUrl:      "http://example.com/product5.jpg",
	}
	productlist = append(productlist, p1, p2, p3, p4, p5)
}
