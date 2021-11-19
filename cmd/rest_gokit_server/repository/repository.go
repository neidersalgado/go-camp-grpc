package repository

import (
	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/models"
)

type Repository interface {
	CreateUser(user models.User) error
	GetUser(userID string) (models.User, error)
	ListUsers() ([]models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(userID string) error
}
