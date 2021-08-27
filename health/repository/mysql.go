package repository

import (
	"fmt"
	models "github.com/chiragsamtani/template-pattern/common/types"
	"github.com/jinzhu/gorm"
)

type MySQLHealthRepository struct {
	DB *gorm.DB
}

func NewMySQLHealthRepository(sess *gorm.DB) *MySQLHealthRepository {
	return &MySQLHealthRepository{
		DB: sess,
	}
}

func (r *MySQLHealthRepository) GetProductByProductCode(productCode string) (*models.Products, error) {
	var product models.Products
	if err := r.DB.First(&product, "product_code = ?", productCode).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *MySQLHealthRepository) CreateProduct(dto *models.Products) error {
	result := r.DB.Create(dto)
	if err := result.Error; err != nil {
		fmt.Println("asokao")
		return err
	}
	return nil
}
