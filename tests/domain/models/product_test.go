package models

import (
	"go-tdd/internal/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate_Product(t *testing.T) {
	// Arrange
	input := models.Product{
		Name:  "Product A",
		Price: 100,
	}

	// Act
	err := input.Validate()

	// Assert
	assert.Equal(t, nil, err)
}

func Test_Validate_Product_Name_Is_Empty(t *testing.T) {
	// Arrange
	input := models.Product{
		Name:  "",
		Price: 100,
	}

	// Act
	err := input.Validate()

	// Assert
	assert.Equal(t, models.ErrNameIsRequired, err)
}

func Test_Validate_Product_Price_Is_Zero(t *testing.T) {
	// Arrange
	input := models.Product{
		Name:  "Product A",
		Price: 0,
	}

	// Act
	err := input.Validate()

	// Assert
	assert.Equal(t, models.ErrPriceMustBeGreaterThanZero, err)
}

func Test_Validate_Product_Is_Nill(t *testing.T) {
	// Arrange
	var input *models.Product

	// Act
	err := input.Validate()

	// Assert
	assert.Equal(t, models.ErrProductIsNill, err)
}
