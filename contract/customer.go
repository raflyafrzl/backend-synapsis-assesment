package contract

import (
	"context"
	"synapsis-test-be/entities"
	"synapsis-test-be/model"
)

type CustomerUseCase interface {
	Find() []entities.UserEntity
	Create(payload model.CreateUserModel) entities.UserEntity
	CreateToken(payload model.CreateUserModel) string
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]entities.UserEntity, error)
	Create(ctx context.Context, payload entities.UserEntity) error
	FindOne(ctx context.Context, payload string) (entities.UserEntity, error)
}
