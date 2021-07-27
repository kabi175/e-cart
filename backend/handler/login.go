package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/kabi175/e-cart/backend/model/apperror"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	user, err := praseUser(r)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	token, err := h.us.Login(user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	cookie := http.Cookie{
		Name:    "jwt-auth",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24 * 5),
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	return
}
