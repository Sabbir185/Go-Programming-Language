package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide me a valid product id", 400)
		return
	}

	product, err := h.svr.Get(id)

	if err != nil {
		http.Error(w, "Failed to get product", 500)
		return
	}
	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product not found!")
		return
	}

	utils.SendJSONData(w, product, 200)
}
