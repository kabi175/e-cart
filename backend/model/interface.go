package model

// UserService must implement the UserService interface
type UserService interface {
	Signup(*User) error
	Login(*User) (string, error)
}

type ProductService interface {
	Create(*Product) error
	Delete(ProductID) error
	FindById(ProductID) (*Product, error)
	FindByCategory(string) (*Product, error)
	FindBySeller(SellerID) (*Product, error)
}

type ProductRepository interface {
	Create(*Product) error
	Delete(string) error
	FindById(string) (*Product, error)
	FindByCategory(string) (*Product, error)
	FindBySeller(string) (*Product, error)
}

// UserRepository must implement the UserRepository interface
type UserRepository interface {
	Create(*User) error
	FindByEmail(string) (*User, error)
}
