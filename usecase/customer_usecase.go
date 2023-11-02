package usecase

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"os"
	"synapsis-test-be/contract"
	"synapsis-test-be/entities"
	"synapsis-test-be/model"
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

func (c *customerUseCase) Create(payload model.CreateUserModel) entities.UserEntity {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var result entities.UserEntity

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		utilities.ErrorResponseWeb(err, 500)
	}
	var id string
	id, _ = gonanoid.New(8)
	result = entities.UserEntity{
		Id:       id,
		Username: payload.Username,
		Password: string(hashed),
	}
	err = c.customer.Create(ctx, result)

	if err != nil {
		panic(model.ResponseFailWeb{
			Error:      "fail while creating customer",
			StatusCode: 400,
			Status:     "fail",
		})
	}
	return result
}

func (c *customerUseCase) CreateToken(payload model.CreateUserModel) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var result entities.UserEntity

	result, _ = c.customer.FindOne(ctx, payload.Username)

	if result.Id == "" {
		utilities.ErrorResponseWeb(errors.New("No Data found"), 400)
	}

	var err error

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(payload.Password))

	if err != nil {
		utilities.ErrorResponseWeb(errors.New("Wrong email/password"), 400)
	}
	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":       result.Id,
		"username": result.Username,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	utilities.ErrorResponseWeb(err, 401)

	return tokenString
}
