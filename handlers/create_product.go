package handlers

import (
	"ecommerce/product"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct product.Product

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newProduct)

	newProduct.ID = len(product.Productlist) + 1
	product.Productlist = append(product.Productlist, newProduct)

	utils.HandleJsonResponse(w, newProduct, 201)
}
