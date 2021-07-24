package mocks

import (
	"github.com/kabi175/e-cart/backend/model"
	"github.com/stretchr/testify/mock"
)

type UserService struct {
	mock.Mock
}

func (o *UserService) Login(user *model.User) (string, error) {
	args := o.Called(user)
	return args.String(0), args.Error(1)
}

func (o *UserService) Signup(user *model.User) error {
	args := o.Called(user)
	return args.Error(0)
}
