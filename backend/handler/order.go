package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

func (h *Handler) FetchOrders(w http.ResponseWriter, r *http.Request) {
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

	orderItems, err := h.os.Fetch(user)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	responce := struct {
		Items []*model.OrderItem `json:"items"`
	}{
		Items: orderItems,
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

func (h *Handler) RemoveOrder(w http.ResponseWriter, r *http.Request) {
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
	userID := vars["userID"]
	orderItem := model.OrderItem{
		SellerID:  user.Email,
		UserID:    userID,
		ProductID: productID,
		Units:     1,
	}

	err = h.os.Remove(&orderItem)
	if err != nil {
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
