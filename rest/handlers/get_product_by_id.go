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

	product := database.Get(id)
	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product not found!")
		return
	}
	utils.SendJSONData(w, product, 200)
}
