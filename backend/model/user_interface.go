package model

// UserService must implement the UserService interface
type UserService interface {
	Signup(*User) error
	Login(*User) error
}

// UserRepository must implement the UserRepository interface
type UserRepository interface {
	Create(*User) error
	FindByEmail(string) (*User, error)
}
