package repository

import (
	"database/sql"

	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

type OrderRepository struct {
	db *sql.DB
}

type OrderConfig struct {
	DB *sql.DB
}

func NewOrderRepository(c *OrderConfig) model.OrderRepository {
	return &OrderRepository{
		db: c.DB,
	}
}

func (o *OrderRepository) Add(item *model.OrderItem) error {
	query := `INSERT INTO orders(seller_id,product_id,user_id,units) VALUES($1, $2, $3, $4)`
	_, err := o.db.Query(query, item.SellerID, item.ProductID, item.UserID, item.Units)
	if err != nil {
		if err == sql.ErrNoRows {
			return apperror.NewConflict("order user_id, product_id", item.UserID+", "+item.ProductID)
		}
		return apperror.NewInternal()
	}
	return nil
}
func (o *OrderRepository) Remove(item *model.OrderItem) error {
	query := `DELETE orders WHERE user_id=$1 and product_id = $2`
	_, err := o.db.Query(query, item.UserID, item.ProductID)
	if err != nil {
		if err == sql.ErrNoRows {
			return apperror.NewNotFound("order user_id, product_id", item.UserID+", "+item.ProductID)
		}
		return apperror.NewInternal()
	}
	return nil
}

func (o *OrderRepository) Fetch(seller *model.User) ([]*model.OrderItem, error) {
	query := `SELECT seller_id,product_id,user_id,units FROM orders  WHERE seller_id  = $1`
	result, err := o.db.Query(query, seller.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.NewNotFound("order user_id", seller.Email)
		}
		return nil, apperror.NewInternal()
	}
	var items []*model.OrderItem
	for result.Next() {
		var item *model.OrderItem
		result.Scan(&item.SellerID, &item.ProductID, &item.UserID, &item.Units)
		items = append(items, item)
	}
	return items, nil
}
