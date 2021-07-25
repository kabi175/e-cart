package service

import (
	"bytes"

	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
	"github.com/segmentio/ksuid"
)

type ProductService struct {
	Pr model.ProductRepository
}

type ProductConfig struct {
	Pr model.ProductRepository
}

func NewProductService(c *ProductConfig) model.ProductService {
	return &ProductService{
		Pr: c.Pr,
	}
}

func (o *ProductService) Create(p *model.Product) error {
	uuid := ksuid.New()
	p.Id = model.ProductID(model.ProductID(uuid.Bytes()))
	err := o.Pr.Create(p)
	if err != nil {
		return apperror.NewInternal()
	}
	return nil
}

func (o *ProductService) Delete(pId model.ProductID) error {
	pIdString := bytes.NewBuffer([]byte(pId)).String()
	err := o.Pr.Delete(pIdString)
	if err != nil {
		if err.Error() == "no match" {
			return apperror.NewNotFound("ProductID", pIdString)
		}
		return apperror.NewInternal()
	}
	return nil
}

func (o *ProductService) FindById(pId model.ProductID) (*model.Product, error) {
	pIdString := bytes.NewBuffer([]byte(pId)).String()
	product, err := o.Pr.FindById(pIdString)
	if err != nil {
		return nil, apperror.NewInternal()
	}
	return product, nil
}

func (o *ProductService) FindByCategory(category string) (*model.Product, error) {
	product, err := o.Pr.FindByCategory(category)
	if err != nil {
		return nil, apperror.NewInternal()
	}
	return product, nil
}

func (o *ProductService) FindBySeller(sellerId model.SellerID) (*model.Product, error) {
	sellerIdString := bytes.NewBuffer([]byte(sellerId)).String()
	product, err := o.Pr.FindBySeller(sellerIdString)
	if err != nil {
		return nil, apperror.NewInternal()
	}
	return product, nil
}
