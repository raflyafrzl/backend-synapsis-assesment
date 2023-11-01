package contract

import (
	"context"
	"synapsis-test-be/entities"
)

type CustomerUseCase interface {
	Find() []entities.UserEntity
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]entities.UserEntity, error)
	Create(ctx context.Context, payload entities.UserEntity) (entities.UserEntity, error)
}
