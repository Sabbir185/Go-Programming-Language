package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
	"sync"
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

	var wg sync.WaitGroup
	var p any
	var c int64

	wg.Add(1)
	go func() {
		defer wg.Add(-1)
		products, err := h.svr.List(page, limit)
		if err != nil {
			http.Error(w, "Failed to get products", 500)
			return
		}
		p = products
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		count := h.svr.Count()
		c = count
	}()

	wg.Wait()

	utils.SendPaginatedData(w, p, c, int64(page), int64(limit))
}
