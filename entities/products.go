package entities

type ProductEntity struct {
	Id          string  `json:"id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}
