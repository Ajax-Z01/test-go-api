package _interface

import "github.com/Ajax-Z01/test-go-api~/entity"

type UserInterface interface {
	GetID() int
	GetUsername() string
	GetEmail() string
}

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(id int) (entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	DeleteUser(id int) error
}
