package contract

import (
	"context"
	"synapsis-test-be/entities"
	"synapsis-test-be/model"
)

type ProductUseCase interface {
	FindAll() []entities.ProductEntity
	FindByCategory(category string) []entities.ProductEntity
	AddToCart(payload model.AddCartModel) entities.Cart
	FindByUserId(userId string) []entities.Cart
}

type ProductRepository interface {
	FindAll(ctx context.Context) []entities.ProductEntity
	FindByCategory(ctx context.Context, category string) []entities.ProductEntity
	SaveToCart(ctx context.Context, data entities.Cart) (entities.Cart, error)
	FindOne(ctx context.Context, id string) entities.ProductEntity
	FindByUserId(ctx context.Context, userId string) []entities.Cart
}
