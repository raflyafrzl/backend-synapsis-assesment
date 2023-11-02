package entities

type ProductEntity struct {
	Id          string  `json:"id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

type Cart struct {
	Id        string `json:"id"`
	ProductId string `json:"products_id" gorm:"column:products_id"`
	UserId    string `json:"user_id" gorm:"column:user_id"`
}
