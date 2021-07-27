package handler

import (
	"log"
	"net/http"

	"github.com/kabi175/e-cart/backend/model/apperror"
)

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	user, err := praseUser(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	err = h.us.Signup(user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}
