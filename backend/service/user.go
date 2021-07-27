package service

import (
	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

type UserService struct {
	userRepo model.UserRepository
	ts       model.TokenService
}

type UserConfig struct {
	UserRepo     model.UserRepository
	TokenService model.TokenService
}

func NewUserService(c *UserConfig) model.UserService {
	return &UserService{
		userRepo: c.UserRepo,
		ts:       c.TokenService,
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
	token, err := u.ts.GenerateToken(user)
	return token, err
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
