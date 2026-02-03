package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendError(w, 400, "Product id is missing")
		return
	}

	var product database.Product
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&product)

	product.ID = id

	updatedProduct := database.Update(product)

	if updatedProduct == nil {
		utils.SendError(w, 400, "Product updates failed")
		return
	}

	utils.SendJSONData(w, updatedProduct, http.StatusOK)
}
