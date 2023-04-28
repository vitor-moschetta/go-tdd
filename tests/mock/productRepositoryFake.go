package mock

import "go-tdd/internal/domain/models"

type ProductRepositoryFake struct {
	Products []models.Product
}

func NewProductLocalRepository() *ProductRepositoryFake {
	return &ProductRepositoryFake{
		Products: []models.Product{},
	}
}

func (r *ProductRepositoryFake) Save(product models.Product) error {
	r.Products = append(r.Products, product)
	return nil
}

func (r *ProductRepositoryFake) FindAll() ([]models.Product, error) {
	return r.Products, nil
}
