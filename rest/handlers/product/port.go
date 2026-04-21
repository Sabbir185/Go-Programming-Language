package product

import "ecommerce/domain"

type Service interface {
	Create(product domain.Product) (*domain.Product, error)
	Delete(id int) error
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
}
