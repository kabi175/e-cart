package service

import "github.com/kabi175/e-cart/backend/model"

type OrderService struct {
	or model.OrderRepository
}

type OrderConfig struct {
	Or model.OrderRepository
}

func NewOrderService(c *OrderConfig) model.OrderService {
	return &OrderService{
		or: c.Or,
	}
}

func (o *OrderService) Add(item *model.OrderItem) error {
	err := o.or.Add(item)
	return err
}
func (o *OrderService) Remove(item *model.OrderItem) error {
	err := o.or.Remove(item)
	return err
}
func (o *OrderService) Fetch(user *model.User) ([]*model.OrderItem, error) {
	items, err := o.or.Fetch(user)
	return items, err
}
