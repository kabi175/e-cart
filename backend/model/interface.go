package model

type UserService interface {
	Login(*User) (string, error)
	Signup(*User) error
}

type UserRepository interface {
	Create(*User) error
	FindByEmail(string) (*User, error)
}
