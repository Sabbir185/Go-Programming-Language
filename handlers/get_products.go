package handlers

import (
	"ecommerce/product"
	"ecommerce/utils"
	"net/http"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleJsonResponse(w, product.Productlist, 200)
}
