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

func (p *productRepository) FindAll(ctx context.Context) []entities.ProductEntity {
	var results []entities.ProductEntity

	err := p.db.Table("products").WithContext(ctx).Find(&results).Error

	utilities.ErrorResponseWeb(err, 500)

	return results

}

func (p *productRepository) FindByCategory(ctx context.Context, category string) []entities.ProductEntity {
	var result []entities.ProductEntity

	var err error = p.db.Table("products").WithContext(ctx).Where("category=?", category).Find(&result).Error

	utilities.ErrorResponseWeb(err, 500)

	return result
}

func (p *productRepository) SaveToCart(ctx context.Context, data entities.Cart) (entities.Cart, error) {

	var err error = p.db.Table("shopping_cart").WithContext(ctx).Create(data).Error

	if err != nil {
		return entities.Cart{}, err
	}

	return data, nil

}

func (p *productRepository) FindOne(ctx context.Context, id string) entities.ProductEntity {

	var result entities.ProductEntity
	var err error = p.db.Table("products").WithContext(ctx).Where("id=?", id).Find(&result).Error

	utilities.ErrorResponseWeb(err, 500)

	return result
}

func (p *productRepository) FindByUserId(ctx context.Context, userId string) []entities.Cart {

	var results []entities.Cart
	var err error = p.db.Table("shopping_cart").WithContext(ctx).Where("user_id=?", userId).Find(&results).Error

	utilities.ErrorResponseWeb(err, 500)

	return results

}
