package infrastructure

import (
	"errors"

	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/entities"
)

type UserGRPCRepository struct{}

func NewUserRepository() *UserGRPCRepository {
	return &UserGRPCRepository{}
}

func (ur *UserGRPCRepository) CreateUser(user entities.User) error {
	return errors.New("Not Implemented")
}

func (ur *UserGRPCRepository) DeleteUser(id string) error {
	return errors.New("Not Implemented")
}

func (ur *UserGRPCRepository) GetUser(userId string) (entities.User, error) {
	if userId == "1234" {
		user := entities.User{
			Id:                    "1234",
			EMail:                 "fake@fake.com",
			Name:                  "fakeName",
			AdditionalInformation: "none",
			Parents:               []entities.User{},
		}

		return user, nil
	}
	return entities.User{}, errors.New("Not Implemented")
}

func (ur *UserGRPCRepository) ListUsers() ([]entities.User, error) {
	return []entities.User{}, errors.New("Not Implemented")
}

func (ur *UserGRPCRepository) UpdateUser(user entities.User) error {
	return errors.New("Not Implemented")
}
