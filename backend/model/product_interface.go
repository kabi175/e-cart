package model

type ProductService interface {
	Create(*Product) error
	Delete(string) error
	FindById(string) (*Product, error)
	FindByCategory(string) ([]*Product, error)
	FindBySeller(string) ([]*Product, error)
	FetchByPage(int) ([]*Product, error)
}

type ProductRepository interface {
	Create(*Product) error
	Delete(string) error
	FindById(string) (*Product, error)
	FindByCategory(string) ([]*Product, error)
	FindBySeller(string) ([]*Product, error)
	Fetch(int, int) ([]*Product, error)
}
