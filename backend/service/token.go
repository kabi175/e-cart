package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kabi175/e-cart/backend/model"
	"github.com/kabi175/e-cart/backend/model/apperror"
)

type TokenService struct {
	key string
}

type TokenConfig struct {
	Key string
}

func NewTokenService(c *TokenConfig) model.TokenService {
	return &TokenService{
		key: c.Key,
	}
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

func (t *TokenService) GenerateToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &Claims{
		Username: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.key))

	if err != nil {
		return "", apperror.NewInternal()
	}
	return tokenString, nil
}

func (t *TokenService) ValidateToken(tokenString string) (*model.User, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.key), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid || !tkn.Valid {
			return nil, apperror.NewAuthorization("Invalid Signature in JWT token")
		}
		return nil, apperror.NewInternal()
	}
	if !tkn.Valid {
		return nil, apperror.NewAuthorization("Invalid JWT token")
	}
	// [TODO] Get user from claims
	return nil, nil
}
