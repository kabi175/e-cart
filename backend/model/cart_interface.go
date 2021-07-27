package model

type CartService interface {
	Add(*User, string) error
	Remove(*User, string) error
	UpdateUnits(*User, *CartProduct) error
	Fetch(*User) ([]*Product, error)
	Order(*User) error
}

type CartRepository interface {
	Add(*User, string) error
	Remove(*User, string) error
	UpdateUnits(*User, *CartProduct) error
	Fetch(*User) ([]*Product, error)
	Order(*User) error
}
