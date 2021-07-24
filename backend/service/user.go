package service

import (
	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

type UserService struct {
	userRepo model.UserRepository
}

type Config struct {
	userRepo model.UserRepository
}

func NewUserService(c *Config) model.UserService {
	return &UserService{
		userRepo: c.userRepo,
	}
}

func (u *UserService) Login(user *model.User) (string, error) {
	// search user in DB
	dbUser, err := u.userRepo.FindByEmail(user.Email)
	if err != nil {
		return "", apperror.NewInternal()
	}
	// if user not found in DB
	if dbUser == nil {
		return "", apperror.NewAuthorization("user email not found")
	}

	if dbUser.Password != user.Password {

		return "", apperror.NewAuthorization("incorrect password")
	}

	// generate token [todo]

	return "token", nil
}

func (u *UserService) Signup(user *model.User) error {
	// search user in DB
	dbUser, err := u.userRepo.FindByEmail(user.Email)
	if err != nil {
		return apperror.NewInternal()
	}

	// if user email conflict
	if dbUser != nil {
		return apperror.NewConflict("user email", user.Email)
	}

	// Create user
	err = u.userRepo.Create(user)
	if err != nil {
		return apperror.NewInternal()
	}
	return nil
}
