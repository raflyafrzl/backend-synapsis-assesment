package contract

import (
	"context"
	"synapsis-test-be/entities"
)

type ProductUseCase interface {
	FindAll() []entities.ProductEntity
}

type ProductRepository interface {
	FindAll(ctx context.Context) []entities.ProductEntity
}
