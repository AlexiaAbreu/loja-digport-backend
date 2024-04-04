package model

type ShoppingCart struct {
	ID          string
	UserId      string
	ProductInfo []ProductInfos
	TotalPrice  float64
	TotalAmount int
}

type ProductInfos struct {
	ProductID     string
	ProductAmount int
}
