package model

// OrderService interface defines Methods to be implemented by OrderService
type OrderService interface {
	Remove(*OrderItem) error
	Fetch(*User) ([]*OrderItem, error)
}

// OrderRepository interface defines Methods to be implemented by OrderRepository
type OrderRepository interface {
	Remove(*OrderItem) error
	Fetch(*User) ([]*OrderItem, error)
}
