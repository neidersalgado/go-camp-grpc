package infrastructure

import (
	"errors"

	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/entities"
)

type UserGRPCRepository struct{}

func NewUserRepository() *UserGRPCRepository {
	return &UserGRPCRepository{}
}

func (up *UserGRPCRepository) GetUserByID(userId int) (entities.User, error) {

	if userId == 1234 {
		user := entities.User{
			Id:       1234,
			EMail:    "fake@fake.com",
			Name:     "fakeName",
			LastName: "fakeLasNAme",
		}

		return user, nil
	}
	return entities.User{}, errors.New("Not Implemented")
}

func (up *UserGRPCRepository) Create(user entities.User) error {
	return errors.New("Not Implemented")
}

func (up *UserGRPCRepository) GetAll() ([]entities.User, error) {
	return []entities.User{}, errors.New("Not Implemented")
}
func (up *UserGRPCRepository) Update(user entities.User) error {
	return errors.New("Not Implemented")
}
func (up *UserGRPCRepository) Delete(userId int) error {
	return errors.New("Not Implemented")
}
