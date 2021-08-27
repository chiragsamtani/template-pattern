package repository

import (
	models "github.com/chiragsamtani/template-pattern/common/types"
)

type HealthRepository interface {
	GetProductByProductCode(productCode string) (*models.Products, error)
	CreateProduct(*models.Products) error
}
