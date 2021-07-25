package mocks

import (
	"github.com/kabi175/e-cart/backend/model"
	"github.com/stretchr/testify/mock"
)

type ProductRepository struct {
	mock.Mock
}

func (o *ProductRepository) Create(p *model.Product) error {
	args := o.Called(p)
	return args.Error(0)
}
func (o *ProductRepository) Delete(pId model.ProductID) error {
	args := o.Called(pId)
	return args.Error(0)
}
func (o *ProductRepository) FindById(pId model.ProductID) (*model.Product, error) {
	args := o.Called(pId)
	arg0 := args.Get(0)
	arg1 := args.Error(1)
	if arg0 == nil {
		return nil, arg1
	}
	return arg0.(*model.Product), arg1
}
func (o *ProductRepository) FindByCategory(category string) (*model.Product, error) {
	args := o.Called(category)
	arg0 := args.Get(0)
	arg1 := args.Error(1)
	if arg0 == nil {
		return nil, arg1
	}
	return arg0.(*model.Product), arg1
}

func (o *ProductRepository) FindBySeller(sellerId model.SellerID) (*model.Product, error) {
	args := o.Called(sellerId)
	arg0 := args.Get(0)
	arg1 := args.Error(1)
	if arg0 == nil {
		return nil, arg1
	}
	return arg0.(*model.Product), arg1
}
