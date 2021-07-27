package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func shutdown(srv *http.Server) {
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, os.Interrupt)
	<-closeChan
	cxt := context.TODO()
	if err := srv.Shutdown(cxt); err != nil {
		log.Printf("Failed to Shutdown: %S", err)
	}
	log.Println("Shuting down server...")
}
