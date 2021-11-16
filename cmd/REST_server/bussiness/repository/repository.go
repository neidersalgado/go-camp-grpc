package repository

import (
	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/entities"
)

type Repository interface {
	CreateUser(user entities.User) error
	GetUser(id string) (entities.User, error)
	ListUsers() ([]entities.User, error)
	UpdateUser(user entities.User) error
	DeleteUser(id string) error
}
