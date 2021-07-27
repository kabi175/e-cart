package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kabi175/e-cart/backend/handler"
)

func main() {
	router := mux.NewRouter()
	handler.NewHandler(&handler.Config{Router: router})

	server := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown(server)

	log.Fatal(server.ListenAndServe())
}
