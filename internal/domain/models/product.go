package models

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNameIsRequired             = errors.New("Name is required")
	ErrPriceMustBeGreaterThanZero = errors.New("Price must be greater than 0")
	ErrProductIsNill              = errors.New("Product is nil")
)

type Product struct {
	ID    string
	Name  string
	Price float32
}

func NewProduct(name string, price float32) Product {
	return Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}

func (p *Product) Validate() error {
	if p == nil {
		return ErrProductIsNill
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price <= 0 {
		return ErrPriceMustBeGreaterThanZero
	}
	return nil
}
