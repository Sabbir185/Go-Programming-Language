package main

import (
	"ecommerce/database"
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func main() {
	// router
	mux := http.NewServeMux()

	// advanced route and middleware
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProductHandler))
	mux.Handle("POST /products/create", http.HandlerFunc(handlers.CreateProductHandler))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	p1 := database.Product{
		ID:          1,
		Title:       "Product 1",
		Price:       19.99,
		Description: "This is the first product.",
		ImgUrl:      "http://example.com/product1.jpg",
	}
	p2 := database.Product{
		ID:          2,
		Title:       "Product 2",
		Price:       29.99,
		Description: "This is the second product.",
		ImgUrl:      "http://example.com/product2.jpg",
	}
	p3 := database.Product{
		ID:          3,
		Title:       "Product 3",
		Price:       39.99,
		Description: "This is the third product.",
		ImgUrl:      "http://example.com/product3.jpg",
	}
	p4 := database.Product{
		ID:          4,
		Title:       "Product 4",
		Price:       49.99,
		Description: "This is the fourth product.",
		ImgUrl:      "http://example.com/product4.jpg",
	}
	p5 := database.Product{
		ID:          5,
		Title:       "Product 5",
		Price:       59.99,
		Description: "This is the fifth product.",
		ImgUrl:      "http://example.com/product5.jpg",
	}
	database.Productlist = append(database.Productlist, p1, p2, p3, p4, p5)
}
