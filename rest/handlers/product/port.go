package product

import "ecommerce/domain"

type Service interface {
	Create(product domain.Product) (*domain.Product, error)
	Delete(id int) error
	Get(id int) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Count()(int64)
	Update(product domain.Product) (*domain.Product, error)
}
