package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 10 {
		limit = 10
	}

	products, err := h.svr.List(page, limit)
	if err != nil {
		http.Error(w, "Failed to get products", 500)
		return
	}

	count := h.svr.Count()

	utils.SendPaginatedData(w, products, count, int64(page), int64(limit))
}
