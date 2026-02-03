package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"log"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Please provide me a valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)
	utils.SendJSONData(w, createdProduct, 201)
}
