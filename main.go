package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
	"synapsis-test-be/contract"
	"synapsis-test-be/database"
	"synapsis-test-be/delivery"
	"synapsis-test-be/middlewares"
	"synapsis-test-be/repository"
	"synapsis-test-be/usecase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var db *gorm.DB = database.NewInitDB()

	var customerRepository contract.CustomerRepository = repository.NewCustomerRepository(db)
	var customerService contract.CustomerUseCase = usecase.NewCustomerUseCase(&customerRepository)
	var customerController *delivery.CustomerController = delivery.NewCustomerController(&customerService)

	r := chi.NewRouter()
	r.Use(middlewares.RecoveryMiddleware)
	r.Route("/api/customers", customerController.Route)

	http.ListenAndServe(":3000", r)
}
