package product

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"log"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Please provide me a valid json", 400)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Failed to create product", 500)
		return
	}

	utils.SendJSONData(w, createdProduct, 201)
}
