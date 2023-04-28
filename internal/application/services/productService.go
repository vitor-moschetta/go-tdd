package services

import (
	"go-tdd/internal/domain/interfaces"
	"go-tdd/internal/domain/models"
)

type ProductService struct {
	Repository interfaces.IProductRepository
	Broker     interfaces.IBroker
}

func NewProductService(Repository interfaces.IProductRepository, broker interfaces.IBroker) *ProductService {
	return &ProductService{
		Repository: Repository,
		Broker:     broker,
	}
}

func (s *ProductService) Create(input models.Product) error {
	// Validate input
	err := input.Validate()
	if err != nil {
		return err
	}

	// Save product
	err = s.Repository.Save(input)
	if err != nil {
		return err
	}

	// Send Async Notification
	go s.Broker.Publish("product.created", input)

	return nil
}
