package model

// User struct defines User-domain and json representation
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Product struct defines Product-domain and json representation
type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	SellerID    string `json:"sellerID"`
	Stock       int64  `json:"stock"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

// Order struct defines Seller Order-domain and json representation
type Order struct {
	SellerID  string `json:"sellerID"`
	UserID    string `json:"userID"`
	ProductID string `json:"productID"`
	Units     int    `json:"units"`
}

// Cart struct defines User Cart-domain and json representation
type Cart struct {
	UserID    string `json:"userID"`
	ProductID string `json:"productID"`
	Units     int    `json:"units"`
}
