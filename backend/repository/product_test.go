package repository

import (
	"testing"

	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db, err := connect()
	assert.NoError(t, err)

	pr := NewProductRepository(&ProductConfig{
		DB: db,
	})

	t.Run("Success", func(t *testing.T) {
		product := model.Product{
			Id:       ksuid.New().String(),
			Name:     "t-shirt",
			Stock:    100,
			SellerID: ksuid.New().String(),
			Category: "fashion",
		}
		err = pr.Create(&product)
		assert.NoError(t, err)
	})
}

func TestDeleteProduct(t *testing.T) {
	db, err := connect()
	assert.NoError(t, err)

	pr := NewProductRepository(&ProductConfig{
		DB: db,
	})

	t.Run("Not found error", func(t *testing.T) {
		err = pr.Delete(ksuid.New().String())
		assert.Equal(t, 404, apperror.Status(err))
	})
}
