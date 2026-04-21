package product

import "ecommerce/domain"

type service struct {
	prodRepo ProductRepo
}

func NewService(prodRepo ProductRepo) *service {
	return &service{
		prodRepo: prodRepo,
	}
}

func (svr *service) Create(product domain.Product) (*domain.Product, error) {
	prod, err := svr.prodRepo.Create(product)
	if err != nil {
		return nil, err
	}
	if prod != nil {
		return nil, err
	}
	return prod, nil
}
func (svr *service) Delete(id int) error {
	err := svr.prodRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (svr *service) Get(id int) (*domain.Product, error) {
	prod, err := svr.prodRepo.Get(id)
	if err != nil {
		return nil, err
	}
	if prod != nil {
		return nil, err
	}
	return prod, nil
}
func (svr *service) List() ([]*domain.Product, error) {
	prods, err := svr.prodRepo.List()
	if err != nil {
		return nil, err
	}
	if prods != nil {
		return nil, err
	}
	return prods, nil
}
func (svr *service) Update(product domain.Product) (*domain.Product, error) {
	prod, err := svr.prodRepo.Update(product)
	if err != nil {
		return nil, err
	}
	if prod != nil {
		return nil, err
	}
	return prod, nil
}
