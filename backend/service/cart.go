package service

import "github.com/kabi175/e-cart/backend/model"

type CartService struct {
	cr model.CartRepository
}

type CartCongig struct {
	Cr model.CartRepository
}

func NewCartService(c *CartCongig) model.CartService {
	return &CartService{
		cr: c.Cr,
	}
}

func (c *CartService) Add(item *model.CartItem) error {
	err := c.cr.Add(item)
	return err
}

func (c *CartService) Remove(item *model.CartItem) error {
	err := c.cr.Remove(item)
	return err
}

func (c *CartService) UpdateUnits(item *model.CartItem) error {
	err := c.cr.UpdateUnits(item)
	return err
}

func (c *CartService) Fetch(user *model.User) ([]*model.CartItem, error) {
	items, err := c.cr.Fetch(user)
	return items, err
}

func (c *CartService) EmptyCart(user *model.User) error {
	err := c.cr.EmptyCart(user)
	return err
}
