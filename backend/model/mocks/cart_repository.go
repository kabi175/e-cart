package mocks

import (
	"github.com/kabi175/e-cart/backend/model"
	"github.com/stretchr/testify/mock"
)

type CartRepository struct {
	mock.Mock
}

func (o *CartRepository) Add(c *model.CartItem) error {
	args := o.Called(c)
	return args.Error(0)
}

func (o *CartRepository) Remove(c *model.CartItem) error {
	args := o.Called(c)
	return args.Error(0)
}

func (o *CartRepository) UpdateUnits(c *model.CartItem) error {
	args := o.Called(c)
	return args.Error(0)
}

func (o *CartRepository) Fetch(u *model.User) ([]*model.CartItem, error) {
	args := o.Called(u)
	arg0 := args.Get(0)
	arg1 := args.Error(1)
	return arg0.([]*model.CartItem), arg1
}

func (o *CartRepository) PlaceOrder(u *model.User) error {
	args := o.Called(u)
	return args.Error(0)
}
