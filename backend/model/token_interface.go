package model

type TokenService interface {
	GenerateToken(*User) (string, error)
	ValidateToken(string) (*User, error)
}
