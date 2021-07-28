package mocks

import (
	"github.com/kabi175/e-cart/backend/model"
	"github.com/stretchr/testify/mock"
)

type OrderRepository struct {
	mock.Mock
}

func (o *OrderRepository) Remove(order *model.OrderItem) error {
	args := o.Called(order)
	return args.Error(0)
}

func (o *OrderRepository) Fetch(u *model.User) ([]*model.OrderItem, error) {
	args := o.Called(u)
	arg0 := args.Get(0)
	arg1 := args.Error(1)
	return arg0.([]*model.OrderItem), arg1
}
