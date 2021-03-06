package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kabi175/e-cart/backend/model"
)

type Handler struct {
	us model.UserService
	ps model.ProductService
	cs model.CartService
	os model.OrderService
	ts model.TokenService
}

type Config struct {
	Router *mux.Router
	Us     model.UserService
	Ps     model.ProductService
	Cs     model.CartService
	Os     model.OrderService
	Ts     model.TokenService
}

func NewHandler(c *Config) {
	// Handler struct
	h := Handler{
		us: c.Us,
		ps: c.Ps,
		cs: c.Cs,
		os: c.Os,
		ts: c.Ts,
	}

	// CROSS orgin resource middleware
	c.Router.Use(crossOriginMiddleware)

	/* Subrouters */
	root := c.Router.PathPrefix("/api").Subrouter()

	//user endpoint subrouter
	user := root.PathPrefix("/user").
		Subrouter()

	user.Use(crossOriginMiddleware)

	//product endpoint subrouter
	product := root.PathPrefix("/product").Subrouter()

	//cart endpoint subrouter
	cart := root.
		PathPrefix("/cart").
		Subrouter()

	//orders endpoint subrouter
	order := root.
		PathPrefix("/order").
		HeadersRegexp("Content-Type", "application/json").
		Subrouter()

	// auth endpoint
	user.HandleFunc("/login", h.Login)
	user.HandleFunc("/signup", h.Signup)

	// products endpoint
	product.HandleFunc("/{id}", h.FindByIdProduct).Methods("GET")                      // get product by ID
	product.HandleFunc("/page/{page}", h.FetchByPageProduct).Methods("GET")            // get product by page
	product.HandleFunc("/seller/{id}", h.FindBySellerProduct).Methods("GET")           // get product by seller
	product.HandleFunc("/category/{category}", h.FindByCategoryProduct).Methods("GET") // get product by category
	product.HandleFunc("/", h.CreateProduct).Methods("POST")                           // create porduct handler
	product.HandleFunc("/{id}", h.DeleteProduct).Methods("DELETE")                     // delete product handler

	// user-cart endpoint
	cart.HandleFunc("/", h.FetchCartItem).Methods("GET")                  // fetch cart products
	cart.HandleFunc("/{productID}", h.AddCartItem).Methods("POST")        // add products to cart
	cart.HandleFunc("/{productID}", h.RemoveCartItem).Methods("DELETE")   // delete product from cart
	cart.HandleFunc("/{productID}", h.UpdateCartItemUnits).Methods("PUT") // update product unit from cart
	cart.HandleFunc("/checkout", h.PlaceOrder).Methods("POST")            // place orders

	// seller-orders endpoint
	order.HandleFunc("/", h.FetchOrders).Methods("GET")                           // fetch orders products
	order.HandleFunc("/{productid}/{productid}", h.RemoveOrder).Methods("DELETE") // delete product from ordes
	root.HandleFunc("/hello", h.Hello)
}

func (Handler) Hello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello endpoint")
	hello := struct {
		Msg string `json:"msg"`
	}{Msg: "hello"}
	body, err := json.Marshal(hello)
	if err != nil {
		log.Println(err)
	}
	w.Write([]byte(body))
}
