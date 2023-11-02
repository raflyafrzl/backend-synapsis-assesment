package delivery

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"synapsis-test-be/contract"
	"synapsis-test-be/entities"
	"synapsis-test-be/model"
	"synapsis-test-be/utilities"
)

type CustomerController struct {
	customer contract.CustomerUseCase
}

func NewCustomerController(c *contract.CustomerUseCase) *CustomerController {
	return &CustomerController{
		customer: *c,
	}
}

func (c *CustomerController) Route(r chi.Router) {
	r.Get("/", c.Get)
	r.Post("/register", c.Create)
	r.Post("/login", c.Login)
}

func (c *CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []entities.UserEntity = c.customer.Find()

	var rawResponse model.ResponseWebSuccess = model.ResponseWebSuccess{
		StatusCode: 200,
		Status:     "success",
		Message:    "success retrieved all customer data",
		Data:       results,
	}
	response, _ := json.Marshal(rawResponse)

	w.WriteHeader(200)
	w.Write(response)

}

func (c *CustomerController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	utilities.ErrorResponseWeb(err, 400)

	var payload model.CreateUserModel
	_ = json.Unmarshal(body, &payload)

	var result entities.UserEntity
	result = c.customer.Create(payload)

	var rawResponse model.ResponseWebSuccess = model.ResponseWebSuccess{
		StatusCode: 201,
		Status:     "success",
		Message:    "a customer data has been successfully created",
		Data:       result,
	}

	response, err := json.Marshal(rawResponse)

	w.WriteHeader(201)
	w.Write(response)
}

func (c *CustomerController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	utilities.ErrorResponseWeb(err, 400)

	var payload model.CreateUserModel
	_ = json.Unmarshal(body, &payload)

	var token string = c.customer.CreateToken(payload)
	var rawResponse model.ResponseWebSuccess = model.ResponseWebSuccess{
		StatusCode: 201,
		Status:     "success",
		Message:    "Token has been created",
		Data:       token,
	}

	response, _ := json.Marshal(rawResponse)

	w.Write(response)

}
