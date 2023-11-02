package model

type AddCartModel struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

type ProductIdModel struct {
	ProductId string `json:"product_id"`
}
