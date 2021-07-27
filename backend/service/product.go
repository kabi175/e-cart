package service

import (
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
	p.Id = ksuid.New().String()
	err := o.Pr.Create(p)
	if err != nil {
		return apperror.NewInternal()
	}
	return nil
}

func (o *ProductService) Delete(pId string) error {
	err := o.Pr.Delete(pId)
	return err
}

func (o *ProductService) FindById(pId string) (*model.Product, error) {
	product, err := o.Pr.FindById(pId)
	if err != nil {
		return nil, apperror.NewInternal()
	}
	return product, nil
}

func (o *ProductService) FindByCategory(category string) ([]*model.Product, error) {
	products, err := o.Pr.FindByCategory(category)
	if err != nil {
		return nil, apperror.NewInternal()
	}
	return products, nil
}

func (o *ProductService) FindBySeller(sellerId string) ([]*model.Product, error) {
	products, err := o.Pr.FindBySeller(sellerId)
	if err != nil {
		return nil, apperror.NewInternal()
	}
	return products, nil
}

func (o *ProductService) FetchByPage(page int) ([]*model.Product, error) {
	maxProductQueryCount := 30
	limit := page * maxProductQueryCount
	products, err := o.Pr.Fetch(limit, (page-1)*maxProductQueryCount)
	if err != nil {
		return nil, apperror.NewInternal()
	}
	return products, nil
}
