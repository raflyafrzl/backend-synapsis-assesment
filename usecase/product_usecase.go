package usecase

import (
	"context"
	"errors"
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"synapsis-test-be/contract"
	"synapsis-test-be/entities"
	"synapsis-test-be/model"
	"synapsis-test-be/utilities"
	"time"
)

type productUseCase struct {
	product contract.ProductRepository
}

func NewProductUseCase(product *contract.ProductRepository) contract.ProductUseCase {

	return &productUseCase{product: *product}
}
func (p productUseCase) FindAll() []entities.ProductEntity {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return p.product.FindAll(ctx)
}

func (p *productUseCase) FindByCategory(category string) []entities.ProductEntity {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	fmt.Print(category)

	return p.product.FindByCategory(ctx, category)
}

func (p *productUseCase) AddToCart(payload model.AddCartModel) entities.Cart {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	result := p.product.FindOne(ctx, payload.ProductId)

	if result.Id == "" {
		utilities.ErrorResponseWeb(errors.New("no product found "), 400)
	}

	id, err := gonanoid.New(8)
	utilities.ErrorResponseWeb(err, 500)

	var data entities.Cart = entities.Cart{
		Id:        id,
		ProductId: payload.ProductId,
		UserId:    payload.UserId,
	}
	var savedCart entities.Cart
	savedCart, err = p.product.SaveToCart(ctx, data)

	utilities.ErrorResponseWeb(err, 400)

	return savedCart

}

func (p *productUseCase) FindByUserId(userId string) []entities.Cart {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	return p.product.FindByUserId(ctx, userId)
}
