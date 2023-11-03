package delivery

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"synapsis-test-be/contract"
	"synapsis-test-be/entities"
	"synapsis-test-be/middlewares"
	"synapsis-test-be/model"
	"synapsis-test-be/utilities"
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

	r.Group(func(r chi.Router) {
		r.Get("/{category}", p.Get)
		r.With(middlewares.AuthMiddleware).Delete("/", p.DeleteProductInCart)
		r.With(middlewares.AuthMiddleware).Post("/cart", p.AddToCart)
		r.With(middlewares.AuthMiddleware).Get("/list", p.ListCart)
		r.With(middlewares.AuthMiddleware).Get("/checkout", p.Checkout)
	})

}

func (p *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category string = chi.URLParam(r, "category")

	var results []entities.ProductEntity = p.product.FindByCategory(category)

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

func (p *ProductController) AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	var productId model.ProductIdModel
	_ = json.Unmarshal(body, &productId)
	utilities.ErrorResponseWeb(err, 400)

	var raw []byte
	var auth model.AuthUserModel
	var requestCtx context.Context = r.Context()

	raw, _ = json.Marshal(requestCtx.Value("auth"))
	_ = json.Unmarshal(raw, &auth)

	var payload model.AddCartModel = model.AddCartModel{
		ProductId: productId.ProductId,
		UserId:    auth.Id,
	}

	var result entities.Cart = p.product.AddToCart(payload)

	var rawResponse model.ResponseWebSuccess = model.ResponseWebSuccess{
		StatusCode: 200,
		Status:     "success",
		Message:    "product has been successfully added to shopping cart",
		Data:       result,
	}

	response, _ := json.Marshal(rawResponse)

	w.Write(response)
}

func (p *ProductController) ListCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var raw []byte
	var auth model.AuthUserModel
	var requestCtx context.Context = r.Context()

	raw, _ = json.Marshal(requestCtx.Value("auth"))
	_ = json.Unmarshal(raw, &auth)

	var result []entities.Cart

	result = p.product.FindByUserId(auth.Id)

	var rawResponse model.ResponseWebSuccess = model.ResponseWebSuccess{
		StatusCode: 200,
		Status:     "success",
		Message:    "Success retrieved all product in cart",
		Data:       result,
	}
	response, _ := json.Marshal(rawResponse)
	w.Write(response)

}

func (p *ProductController) DeleteProductInCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var raw []byte
	var auth model.AuthUserModel
	var requestCtx context.Context = r.Context()
	var id string = r.URL.Query().Get("id")

	raw, _ = json.Marshal(requestCtx.Value("auth"))
	_ = json.Unmarshal(raw, &auth)
	p.product.DeleteOneProduct(id)
	var rawResponse model.ResponseWebSuccess = model.ResponseWebSuccess{
		StatusCode: 200,
		Status:     "success",
		Message:    "Success deleted all product in cart",
		Data:       []string{},
	}
	response, _ := json.Marshal(rawResponse)
	w.Write(response)
}

func (p *ProductController) Checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var raw []byte
	var auth model.AuthUserModel
	var requestCtx context.Context = r.Context()

	raw, _ = json.Marshal(requestCtx.Value("auth"))
	_ = json.Unmarshal(raw, &auth)
	p.product.DeleteAllCartById(auth.Id)
	var rawResponse model.ResponseWebSuccess = model.ResponseWebSuccess{
		StatusCode: 200,
		Status:     "success",
		Message:    "Success checkouts product in cart",
		Data:       []string{},
	}
	response, _ := json.Marshal(rawResponse)
	w.Write(response)
}
