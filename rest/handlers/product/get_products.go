package product

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svr.List()
	if err != nil {
		http.Error(w, "Failed to get products", 500)
		return
	}
	utils.SendJSONData(w, products, 200)
}
