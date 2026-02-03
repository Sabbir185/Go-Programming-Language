package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide me a valid product id", 400)
		return
	}

	for _, product := range database.Productlist {
		if product.ID == id {
			utils.HandleJsonResponse(w, product, http.StatusOK)
			return
		}
	}

	http.Error(w, "Product not found", 404)
}
