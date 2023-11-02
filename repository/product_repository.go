package repository

import (
	"context"
	"gorm.io/gorm"
	"synapsis-test-be/contract"
	"synapsis-test-be/entities"
	"synapsis-test-be/utilities"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) contract.ProductRepository {
	return &productRepository{db: db}
}

func (p productRepository) FindAll(ctx context.Context) []entities.ProductEntity {
	//TODO implement me

	var results []entities.ProductEntity

	err := p.db.Table("products").WithContext(ctx).Find(&results).Error

	utilities.ErrorResponseWeb(err, 500)

	return results

}
