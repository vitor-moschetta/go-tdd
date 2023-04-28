package interfaces

import "go-tdd/internal/domain/models"

type IProductRepository interface {
	Save(product models.Product) error
	FindAll() ([]models.Product, error)
}
