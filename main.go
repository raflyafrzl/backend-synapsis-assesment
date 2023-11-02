package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
	"os"
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

	var db *gorm.DB = database.NewInitDB(os.Getenv("DATA_SOURCE_NAME"))

	var customerRepository contract.CustomerRepository = repository.NewCustomerRepository(db)
	var customerService contract.CustomerUseCase = usecase.NewCustomerUseCase(&customerRepository)
	var customerController *delivery.CustomerController = delivery.NewCustomerController(&customerService)

	var productController *delivery.ProductController = delivery.NewProductController()

	r := chi.NewRouter()
	r.Use(middlewares.RecoveryMiddleware)
	r.Route("/api/customers", customerController.Route)
	r.Route("/api/products", productController.Route)
	fmt.Print("Server running on port ", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), r)
}
