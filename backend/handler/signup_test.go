package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
	"github.com/kabi175/e-cart/backend/model/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	userService := new(mocks.UserService)
	router := mux.NewRouter()
	NewHandler(&Config{
		Router: router,
		Us:     userService,
	})

	t.Run("success", func(t *testing.T) {
		user := model.User{
			Email:    "kabi@gmail.com",
			Password: "password",
			Username: "kabi",
		}
		reqBody, err := json.Marshal(user)
		assert.NoError(t, err)
		req, err := http.NewRequest("GET", "/api/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		userService.On("Signup", &user).Return(nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("email conflict", func(t *testing.T) {
		user := model.User{
			Email:    "kabi@gmail.com",
			Password: "password",
		}
		reqBody, err := json.Marshal(user)
		assert.NoError(t, err)
		req, err := http.NewRequest("GET", "/api/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		userService.On("Signup", &user).Return(apperror.NewConflict("user email", "kabi@gmail.com"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
	})
}
