package handler

import (
	"github.com/gorilla/mux"
	"github.com/kabi175/e-cart/backend/model"
)

type Handler struct {
	us model.UserService
	ps model.ProductService
}

type Config struct {
	Router *mux.Router
	Us     model.UserService
	Ps     model.ProductService
}

func NewHandler(c *Config) {
	// Handler struct
	h := Handler{
		us: c.Us,
		ps: c.Ps,
	}

	// CROSS orgin resource middleware
	c.Router.Use(crossOriginMiddleware)

	// Subrouter -> filter json content
	root := c.Router.
		PathPrefix("/api").
		Subrouter()

	auth := root.
		PathPrefix("/auth").
		HeadersRegexp("Content-Type", "application/json").
		Subrouter()

	// user endpoint
	auth.HandleFunc("/login", h.Login)
	auth.HandleFunc("/signup", h.Signup)

	// products endpoint
	product := root.PathPrefix("/product").Subrouter()
	product.HandleFunc("/{id}", h.FindByIdProduct).Methods("GET")                      // get product by ID
	product.HandleFunc("/page/{page}", h.FetchByPageProduct).Methods("GET")            // get product by page
	product.HandleFunc("/seller/{id}", h.FindBySellerProduct).Methods("GET")           // get product by seller
	product.HandleFunc("/category/{category}", h.FindByCategoryProduct).Methods("GET") // get product by category
	product.HandleFunc("/", h.CreateProduct).Methods("POST")                           // create porduct handler
	product.HandleFunc("/{id}", h.DeleteProduct).Methods("DELETE")                     // delete product handler

	// user endpoint
	user := root.
		PathPrefix("/user").
		HeadersRegexp("Content-Type", "application/json").
		Subrouter()

	user.HandleFunc("/cart", h.Login).Methods("GET")    // get cart products
	user.HandleFunc("/cart", h.Login).Methods("POST")   // add products to cart
	user.HandleFunc("/cart", h.Login).Methods("DELETE") // delete products from cart
}
