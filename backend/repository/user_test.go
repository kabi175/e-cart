package repository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/kabi175/e-cart/backend/model"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "ecart"
)

func TestCreate(t *testing.T) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	assert.NoError(t, err)

	t.Run("Success", func(t *testing.T) {
		user := model.User{
			Email:    "kabi@gmail.com",
			Password: "password",
			Username: "kabi",
		}
		userRepo := NewUserRepository(&Config{
			DB: db,
		})
		got := userRepo.Create(&user)
		assert.NoError(t, got)
	})
}
