package services

import (
	"go-tdd/internal/application/services"
	"go-tdd/internal/domain/models"
	"go-tdd/tests/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Product(t *testing.T) {
	// Arrange
	input := models.Product{
		Name:  "Product A",
		Price: 100,
	}
	repository := mock.NewProductLocalRepository()
	broker := mock.NewBrokerFake()

	// Act
	service := services.NewProductService(repository, broker)
	err := service.Create(input)

	// Assert
	assert.Equal(t, nil, err)
}
