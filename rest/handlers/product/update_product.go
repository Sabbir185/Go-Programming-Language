package product

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendError(w, 400, "Product id is missing")
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&req)

	updatedProduct, err := h.productRepo.Update(repo.Product{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})

	if err != nil {
		utils.SendError(w, 500, "Failed to update product")
		return
	}
	if updatedProduct == nil {
		utils.SendError(w, 400, "Product updates failed")
		return
	}

	utils.SendJSONData(w, "Product updated successfully!", http.StatusOK)
}
