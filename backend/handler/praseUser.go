package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

func validate(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		return apperror.NewBadRequest(err.Error())
	}
	return nil
}

func praseUser(r *http.Request) (*model.User, error) {
	user := struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Username string `json:"username"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, apperror.NewBadRequest(err.Error())
	}
	log.Println(err)
	err = validate(user)
	if err != nil {
		return nil, apperror.NewBadRequest(err.Error())
	}

	return &model.User{Email: user.Email, Password: user.Password, Username: user.Username}, nil
}
