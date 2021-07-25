package model

type ProductID []byte
type SellerID []byte

type Product struct {
	Id       ProductID `json:"id"`
	Name     string    `json:"name"`
	SellerID SellerID  `json:"sellerId"`
	Stock    int64     `json:"stock"`
	Category string    `json:"category"`
}
