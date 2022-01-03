package model

// CartService interface defines Methods to be implemented by CartService
type CartService interface {
	Add(*CartItem) error
	Remove(*CartItem) error
	UpdateUnits(*CartItem) error
	Fetch(*User) ([]*CartItem, error)
	EmptyCart(*User) error
}

// CartRepository interface defines Methods to be implemented by CartRepository
type CartRepository interface {
	Add(*CartItem) error
	Remove(*CartItem) error
	UpdateUnits(*CartItem) error
	Fetch(*User) ([]*CartItem, error)
	EmptyCart(*User) error
}
