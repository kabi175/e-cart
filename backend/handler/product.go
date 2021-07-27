package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	p := struct {
		Name     string `json:"name" validate:"required"`
		SellerID string `json:"sellerId" validate:"required"`
		Stock    int64  `json:"stock" validate:"required"`
		Category string `json:"category" validate:"required"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		err = apperror.NewInternal()
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	err = validate(&p)
	if err != nil {
		log.Println(err)
		err = apperror.NewInternal()
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	product := model.Product{
		Name:     p.Name,
		SellerID: p.SellerID,
		Stock:    p.Stock,
		Category: p.Category,
	}
	err = h.ps.Create(&product)
	if err != nil {
		log.Println(err)
		err = apperror.NewInternal()
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]
	err := h.ps.Delete(productID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) FindByCategoryProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	products, err := h.ps.FindByCategory(category)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	responce := struct {
		Product []*model.Product `json:"products"`
	}{
		Product: products,
	}

	responseBody, err := json.Marshal(responce)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (h *Handler) FindBySellerProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sellerID := vars["id"]

	products, err := h.ps.FindByCategory(sellerID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	responce := struct {
		Products []*model.Product `json:"products"`
	}{
		Products: products,
	}

	responseBody, err := json.Marshal(responce)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
func (h *Handler) FindByIdProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]
	product, err := h.ps.FindById(productID)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	responce := struct {
		Product *model.Product `json:"product"`
	}{
		Product: product,
	}

	responseBody, err := json.Marshal(responce)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
func (h *Handler) FetchByPageProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageNumberString := vars["page"]

	pageNumberInt, err := strconv.Atoi(pageNumberString)
	if err != nil {
		err = apperror.NewBadRequest(err.Error())
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	products, err := h.ps.FetchByPage(pageNumberInt)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	responce := struct {
		Product []*model.Product `json:"product"`
	}{
		Product: products,
	}

	responseBody, err := json.Marshal(responce)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), apperror.Status(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
