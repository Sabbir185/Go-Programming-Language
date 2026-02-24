package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONData(w, database.List(), 200)
}
