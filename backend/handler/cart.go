package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

func (h *Handler) AddCartItem(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt-auth")
	if err != nil {
		err = apperror.NewAuthorization("jwt-auth token not found")
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	user, err := h.ts.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	vars := mux.Vars(r)
	productID := vars["productID"]
	cartItem := model.CartItem{
		UserID:    user.Email,
		ProductID: productID,
		Units:     1,
	}
	err = h.cs.Add(&cartItem)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) RemoveCartItem(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt-auth")
	if err != nil {
		err = apperror.NewAuthorization("jwt-auth token not found")
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	user, err := h.ts.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	vars := mux.Vars(r)
	productID := vars["productID"]
	cartItem := model.CartItem{
		UserID:    user.Email,
		ProductID: productID,
	}
	err = h.cs.Remove(&cartItem)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateCartItemUnits(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt-auth")
	if err != nil {
		err = apperror.NewAuthorization("jwt-auth token not found")
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	user, err := h.ts.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	vars := mux.Vars(r)
	productID := vars["productID"]
	cartItem := model.CartItem{
		UserID:    user.Email,
		ProductID: productID,
	}
	err = h.cs.UpdateUnits(&cartItem)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) FetchCartItem(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt-auth")
	if err != nil {
		err = apperror.NewAuthorization("jwt-auth token not found")
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	user, err := h.ts.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	cartItems, err := h.cs.Fetch(user)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	responce := struct {
		Items []*model.CartItem `json:"items"`
	}{
		Items: cartItems,
	}
	if err != nil {
		err = apperror.NewInternal()
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	responceBody, err := json.Marshal(responce)
	w.WriteHeader(http.StatusOK)
	w.Write(responceBody)
}

func (h *Handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt-auth")
	if err != nil {
		err = apperror.NewAuthorization("jwt-auth token not found")
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	user, err := h.ts.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	items, err := h.cs.Fetch(user)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	for _, cartItem := range items {
		orderItem := model.OrderItem{
			ProductID: cartItem.ProductID,
			UserID:    cartItem.UserID,
		}
		err := h.os.Add(&orderItem)
		if err != nil {
			http.Error(w, err.Error(), apperror.Status(err))
			return
		}
	}

	err = h.cs.EmptyCart(user)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
