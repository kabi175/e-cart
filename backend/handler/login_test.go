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

func TestLogin(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		user := model.User{
			Email:    "kabi@gmail.com",
			Password: "password",
		}
		reqBody, err := json.Marshal(user)
		assert.NoError(t, err)
		userService := new(mocks.UserService)
		userService.On("Login", &user).Return("token", nil)

		router := mux.NewRouter()
		NewHandler(&Config{
			Router: router,
			Us:     userService,
		})

		req, err := http.NewRequest("GET", "/api/login", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("auth error", func(t *testing.T) {
		user := model.User{
			Email:    "kabi@gmail.com",
			Password: "incorrect",
		}
		reqBody, err := json.Marshal(user)
		assert.NoError(t, err)
		userService := new(mocks.UserService)
		userService.On("Login", &user).Return("", apperror.NewAuthorization("incorrect password"))

		router := mux.NewRouter()
		NewHandler(&Config{
			Router: router,
			Us:     userService,
		})

		req, err := http.NewRequest("GET", "/api/login", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("validation error", func(t *testing.T) {
		user := model.User{
			Email:    "kabi",
			Password: "incorrect",
		}
		reqBody, err := json.Marshal(user)
		assert.NoError(t, err)
		userService := new(mocks.UserService)
		userService.On("Login", &user).Return("", nil)

		router := mux.NewRouter()
		NewHandler(&Config{
			Router: router,
			Us:     userService,
		})

		req, err := http.NewRequest("GET", "/api/login", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
