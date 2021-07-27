package main

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/kabi175/e-cart/backend/handler"
	"github.com/kabi175/e-cart/backend/repository"
	"github.com/kabi175/e-cart/backend/service"
)

func inject() *mux.Router {
	// Load Evn
	secretKey := os.Getenv("JWT_SECRET")

	db := pgConnect()

	userRepository := repository.NewUserRepository(&repository.Config{DB: db})

	tokenService := service.NewTokenService(&service.TokenConfig{
		Key: secretKey,
	})

	userService := service.NewUserService(&service.UserConfig{
		UserRepo:     userRepository,
		TokenService: tokenService,
	})

	ProductRepository := repository.NewProductRepository(&repository.ProductConfig{DB: db})

	ProductService := service.NewProductService(&service.ProductConfig{
		Pr: ProductRepository,
	})

	router := mux.NewRouter()
	handler.NewHandler(&handler.Config{
		Router: router,
		Us:     userService,
		Ps:     ProductService,
	})
	return router
}
