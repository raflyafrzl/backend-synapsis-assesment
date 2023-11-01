package repository

import (
	"context"
	"gorm.io/gorm"
	"synapsis-test-be/contract"
	"synapsis-test-be/entities"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) contract.CustomerRepository {
	return &customerRepository{db: db}
}

func (c *customerRepository) FindAll(ctx context.Context) ([]entities.UserEntity, error) {

	var results []entities.UserEntity

	var err error = c.db.Table("customers").WithContext(ctx).Find(&results).Error

	if err != nil {
		return nil, err
	}
	return results, nil

}

func (c *customerRepository) Create(ctx context.Context, payload entities.UserEntity) (entities.UserEntity, error) {
	var err error

	err = c.db.Table("customers").WithContext(ctx).Create(payload).Error

	if err != nil {
		return entities.UserEntity{}, err
	}
	return payload, nil
}
