package mocks

import (
	"github.com/kabi175/e-cart/backend/model"
	"github.com/stretchr/testify/mock"
)

type ProductService struct {
	mock.Mock
}

func (o *ProductService) Create(p *model.Product) error {
	args := o.Called(p)
	return args.Error(0)
}

func (o *ProductService) Delete(pId string) error {
	args := o.Called(pId)
	return args.Error(0)
}

func (o *ProductService) FindById(pId string) (*model.Product, error) {
	args := o.Called(pId)
	arg0 := args.Get(0)
	arg1 := args.Error(1)
	if arg0 == nil {
		return nil, arg1
	}
	return arg0.(*model.Product), arg1
}

func (o *ProductService) FindByCategory(category string) ([]*model.Product, error) {
	args := o.Called(category)
	arg0 := args.Get(0)
	arg1 := args.Error(1)
	if arg0 == nil {
		return nil, arg1
	}
	return arg0.([]*model.Product), arg1
}

func (o *ProductService) FindBySeller(sellerId string) ([]*model.Product, error) {
	args := o.Called(sellerId)
	arg0 := args.Get(0)
	arg1 := args.Error(1)
	if arg0 == nil {
		return nil, arg1
	}
	return arg0.([]*model.Product), arg1
}

func (o *ProductService) Fetch(limit, offset int) ([]*model.Product, error) {
	args := o.Called(limit, offset)
	arg0 := args.Get(0)
	arg1 := args.Error(1)
	if arg0 == nil {
		return nil, arg1
	}
	return arg0.([]*model.Product), arg1
}
