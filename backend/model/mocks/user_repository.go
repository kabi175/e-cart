package mocks

import (
	"github.com/kabi175/e-cart/backend/model"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (o *UserRepository) Create(user *model.User) error {
	args := o.Called(user)
	return args.Error(0)
}

func (o *UserRepository) FindByEmail(email string) (*model.User, error) {
	args := o.Called(email)
	arg0 := args.Get(0)
	arg1 := args.Error(1)

	if arg0 == nil {
		return nil, arg1
	}
	return arg0.(*model.User), arg1
}
