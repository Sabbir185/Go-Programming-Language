package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleJsonResponse(w, database.Productlist, 200)
}
