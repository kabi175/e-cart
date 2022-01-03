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

	productRepository := repository.NewProductRepository(&repository.ProductConfig{DB: db})

	cartRepository := repository.NewCartRepository(&repository.CartConfig{DB: db})

	orderRepository := repository.NewOrderRepository(&repository.OrderConfig{DB: db})

	userService := service.NewUserService(&service.UserConfig{
		UserRepo: userRepository,
	})

	productService := service.NewProductService(&service.ProductConfig{
		Pr: productRepository,
	})

	cartService := service.NewCartService(&service.CartCongig{
		Cr: cartRepository,
	})

	orderService := service.NewOrderService(&service.OrderConfig{
		Or: orderRepository,
	})

	tokenService := service.NewTokenService(&service.TokenConfig{
		Key: secretKey,
	})

	router := mux.NewRouter()
	handler.NewHandler(&handler.Config{
		Router: router,
		Us:     userService,
		Ps:     productService,
		Cs:     cartService,
		Os:     orderService,
		Ts:     tokenService,
	})
	return router
}
