package model

// OrderService interface defines Methods to be implemented by OrderService
type OrderService interface {
	Add(*OrderItem) error
	Remove(*OrderItem) error
	Fetch(*User) ([]*OrderItem, error)
}

// OrderRepository interface defines Methods to be implemented by OrderRepository
type OrderRepository interface {
	Add(*OrderItem) error
	Remove(*OrderItem) error
	Fetch(*User) ([]*OrderItem, error)
}
