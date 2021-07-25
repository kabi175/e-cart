package handler

import (
	"net/http"
)

func crossOriginMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		handler.ServeHTTP(w, r)
	})
}

func authMiddleware(handler http.Handler) http.Handler {
	panic("authMiddleware not implemented")
}
