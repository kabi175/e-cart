package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

func praseUser(r *http.Request) (*model.User, error) {
	user := struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, apperror.NewBadRequest(err.Error())
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		return nil, apperror.NewBadRequest(err.Error())
	}

	return &model.User{Email: user.Email, Password: user.Password}, nil
}
