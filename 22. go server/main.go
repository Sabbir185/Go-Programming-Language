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

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"imgUrl"`
}

var productlist []Product

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Request, not found!", 400)
		return
	}
	handleJsonResponse(w, productlist, 200)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request, not found!", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newProduct)

	newProduct.ID = len(productlist) + 1
	productlist = append(productlist, newProduct)

	handleJsonResponse(w, newProduct, 201)
}

func handleHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func handleJsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux() // router

	// advanced route and middleware
	mux.Handle("GET /hello", http.HandlerFunc(handleHelloHandler))
	mux.Handle("GET /about", http.HandlerFunc(aboutMeHandler))
	mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(getProductHandler)))
	mux.Handle("POST /products/create", corsMiddleware(http.HandlerFunc(createProductHandler)))

	// old style route
	// mux.HandleFunc("/products/create", createProductHandler)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	p1 := Product{
		ID:          1,
		Title:       "Product 1",
		Price:       19.99,
		Description: "This is the first product.",
		ImgUrl:      "http://example.com/product1.jpg",
	}
	p2 := Product{
		ID:          2,
		Title:       "Product 2",
		Price:       29.99,
		Description: "This is the second product.",
		ImgUrl:      "http://example.com/product2.jpg",
	}
	p3 := Product{
		ID:          3,
		Title:       "Product 3",
		Price:       39.99,
		Description: "This is the third product.",
		ImgUrl:      "http://example.com/product3.jpg",
	}
	p4 := Product{
		ID:          4,
		Title:       "Product 4",
		Price:       49.99,
		Description: "This is the fourth product.",
		ImgUrl:      "http://example.com/product4.jpg",
	}
	p5 := Product{
		ID:          5,
		Title:       "Product 5",
		Price:       59.99,
		Description: "This is the fifth product.",
		ImgUrl:      "http://example.com/product5.jpg",
	}
	productlist = append(productlist, p1, p2, p3, p4, p5)
}

// Middleware
func corsMiddleware(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleCors)
}
