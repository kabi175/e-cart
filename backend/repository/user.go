package repository

import (
	"database/sql"
	"log"

	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

type Config struct {
	DB *sql.DB
}

func NewUserRepository(c *Config) model.UserRepository {
	return &UserRepository{
		DB: c.DB,
	}
}

func (u *UserRepository) Create(user *model.User) error {
	query := "INSERT INTO users(email, password, username) VALUES( $1 , $2 , $3 )"

	result, err := u.DB.Exec(query, user.Email, user.Password, user.Username)
	if err != nil {
		log.Println(err)
		return apperror.NewInternal()
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return apperror.NewInternal()
	}

	if rows != 1 {
		log.Println("UserRepository.Create: rows affected not equal to 1")
		return apperror.NewInternal()
	}

	return nil
}
func (u *UserRepository) FindByEmail(email string) (*model.User, error) {
	var (
		Email    string
		Password string
		Username string
	)
	query := `SELECT email,password,username FROM users WHERE email=$1`
	err := u.DB.QueryRow(query, email).Scan(&Email, &Password, &Username)

	// if email not found
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Println(err)
		return nil, apperror.NewInternal()
	}

	// if succusses
	user := model.User{
		Email:    Email,
		Password: Password,
		Username: Username,
	}
	return &user, nil
}
