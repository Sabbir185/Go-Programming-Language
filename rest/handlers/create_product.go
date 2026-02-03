package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newProduct)

	newProduct.ID = len(database.Productlist) + 1
	database.Productlist = append(database.Productlist, newProduct)

	utils.HandleJsonResponse(w, newProduct, 201)
}
