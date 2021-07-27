package model

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	SellerID string `json:"sellerId"`
	Stock    int64  `json:"stock"`
	Category string `json:"category"`
}
