package usecase

import (
	"context"
	"synapsis-test-be/contract"
	"synapsis-test-be/entities"
	"synapsis-test-be/utilities"
	"time"
)

type customerUseCase struct {
	customer contract.CustomerRepository
}

func NewCustomerUseCase(repo *contract.CustomerRepository) contract.CustomerUseCase {
	return &customerUseCase{
		customer: *repo,
	}
}

func (c *customerUseCase) Find() []entities.UserEntity {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var results []entities.UserEntity
	var err error

	results, err = c.customer.FindAll(ctx)
	utilities.ErrorResponseWeb(err, 500)

	return results
}
