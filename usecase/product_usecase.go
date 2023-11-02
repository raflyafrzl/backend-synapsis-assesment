package usecase

import (
	"context"
	"synapsis-test-be/contract"
	"synapsis-test-be/entities"
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
