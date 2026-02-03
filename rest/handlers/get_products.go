package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONData(w, database.List(), 200)
}
