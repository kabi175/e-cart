package service

import (
	"errors"
	"testing"

	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
	"github.com/kabi175/e-cart/backend/model/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		user := model.User{
			Username: "kabi175",
			Email:    "kabilan@gmail.com",
			Password: "password",
			Type:     "user",
		}
		userRepo := new(mocks.UserRepository)
		userRepo.On("FindByEmail", user.Email).Return(nil, nil)
		userRepo.On("Create", &user).Return(nil)
		service := NewUserService(&Config{userRepo: userRepo})
		got := service.Signup(&user)
		assert.Equal(t, nil, got)
	})

	t.Run("email conflict", func(t *testing.T) {
		user := model.User{
			Username: "kabi175",
			Email:    "kabilan@gmail.com",
			Password: "password",
			Type:     "user",
		}
		userRepo := new(mocks.UserRepository)
		userRepo.On("FindByEmail", user.Email).Return(&user, nil)
		userRepo.On("Create", &user).Return(nil)

		service := NewUserService(&Config{userRepo: userRepo})
		got := service.Signup(&user)
		assert.Equal(t, apperror.NewConflict("user email", user.Email), got)
	})

	t.Run("internal error", func(t *testing.T) {
		user := model.User{
			Username: "kabi175",
			Email:    "kabilan@gmail.com",
			Password: "password",
			Type:     "user",
		}
		userRepo := new(mocks.UserRepository)
		userRepo.On("FindByEmail", user.Email).Return(nil, nil)
		userRepo.On("Create", &user).Return(errors.New("failed to reach DB"))

		service := NewUserService(&Config{userRepo: userRepo})
		got := service.Signup(&user)
		assert.Equal(t, apperror.NewInternal(), got)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := model.User{
			Email:    "kabilan@gmail.com",
			Password: "password",
		}
		userRepo := new(mocks.UserRepository)
		userRepo.On("FindByEmail", user.Email).Return(&user, nil)

		service := NewUserService(&Config{userRepo: userRepo})
		token, err := service.Login(&user)
		assert.Equal(t, nil, err, "err must be nil")
		assert.NotEqual(t, "", token, "token should not be empty")
	})

	t.Run("auth error", func(t *testing.T) {
		user := model.User{
			Email:    "kabilan@gmail.com",
			Password: "password",
		}
		dbUser := model.User{
			Email:    "kabilan@gmail.com",
			Password: "notmatch",
		}
		userRepo := new(mocks.UserRepository)
		userRepo.On("FindByEmail", user.Email).Return(&dbUser, nil)

		service := NewUserService(&Config{userRepo: userRepo})
		token, err := service.Login(&user)
		assert.Equal(t, apperror.NewAuthorization("incorrect password"), err, "err must be Auth error")
		assert.Equal(t, "", token, "token should be empty")
	})

	t.Run("user email not-found error", func(t *testing.T) {
		user := model.User{
			Email:    "kabilan@gmail.com",
			Password: "password",
		}
		userRepo := new(mocks.UserRepository)
		userRepo.On("FindByEmail", user.Email).Return(nil, nil)

		service := NewUserService(&Config{userRepo: userRepo})
		token, err := service.Login(&user)
		assert.Equal(t, apperror.NewAuthorization("user email not found"), err, "err must be nil")
		assert.Equal(t, "", token, "token should  be empty")
	})
}
