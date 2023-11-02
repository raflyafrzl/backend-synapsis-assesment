package delivery

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"synapsis-test-be/contract"
	"synapsis-test-be/entities"
	"synapsis-test-be/model"
)

type ProductController struct {
	product contract.ProductUseCase
}

func NewProductController(product *contract.ProductUseCase) *ProductController {

	return &ProductController{
		product: *product,
	}
}

func (p *ProductController) Route(r chi.Router) {
	r.Get("/", p.Get)
}

func (p *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []entities.ProductEntity = p.product.FindAll()

	var rawResponse model.ResponseWebSuccess = model.ResponseWebSuccess{
		StatusCode: 200,
		Status:     "success",
		Message:    "Products has been retrieved successfully",
		Data:       results,
	}

	var body []byte
	body, _ = json.Marshal(rawResponse)

	w.WriteHeader(200)
	w.Write(body)
}
