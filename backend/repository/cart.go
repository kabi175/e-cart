package repository

import (
	"database/sql"

	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

type CartRepository struct {
	db *sql.DB
}

type CartConfig struct {
	DB *sql.DB
}

func NewCartRepository(c *CartConfig) model.CartRepository {
	return &CartRepository{
		db: c.DB,
	}
}

func (c *CartRepository) Add(item *model.CartItem) error {
	query := `INSET INTO cart(user_id,product_it,units) VALUES($1,$2,$3)`
	err := c.db.QueryRow(query, item.UserID, item.ProductID, item.Units).Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return apperror.NewConflict("user_id, product_id", item.UserID+", "+item.ProductID)
		}
		return apperror.NewInternal()
	}
	return nil
}

func (c *CartRepository) Remove(item *model.CartItem) error {
	query := `DELETE cart
						WHERE user_id=$1 and product_id=$2`
	_, err := c.db.Query(query, item.UserID, item.ProductID)
	if err != nil {
		if err == sql.ErrNoRows {
			return apperror.NewNotFound("user_id, product_id", item.UserID+", "+item.ProductID)
		}
		return apperror.NewInternal()
	}
	return nil
}

func (c *CartRepository) UpdateUnits(item *model.CartItem) error {
	query := `UPDATE cart
						SET units=$3
						WHERE user_id=$1 and product_id=$2`
	_, err := c.db.Query(query, item.UserID, item.ProductID, item.Units)
	if err != nil {
		if err == sql.ErrNoRows {
			return apperror.NewNotFound("user_id, product_id", item.UserID+", "+item.ProductID)
		}
		return apperror.NewInternal()
	}
	return nil
}

func (c *CartRepository) Fetch(user *model.User) ([]*model.CartItem, error) {
	query := `SELECT * 
						FROM cart
						WHERE cart.user_id = $1`
	result, err := c.db.Query(query, user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.NewNotFound("user_id", user.Email)
		}
		return nil, apperror.NewInternal()
	}
	var items []*model.CartItem
	for result.Next() {
		var item *model.CartItem
		result.Scan(&item.UserID, &item.ProductID, &item.Units)
		items = append(items, item)
	}
	return items, nil
}

func (c *CartRepository) EmptyCart(user *model.User) error {
	query := `DELETE cart
						WHERE cart.user_id = $1`
	_, err := c.db.Query(query, user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return apperror.NewNotFound("user_id", user.Email)
		}
		return apperror.NewInternal()
	}
	return nil
}
