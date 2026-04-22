package product

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide me a valid product id", 400)
		return
	}

	err = h.svr.Delete(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to delete product", 500)
		return
	}

	utils.SendJSONData(w, "Product deleted successfully", 200)
}
