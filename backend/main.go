package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := inject()
	server := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown(server)

	log.Fatal(server.ListenAndServe())
}
