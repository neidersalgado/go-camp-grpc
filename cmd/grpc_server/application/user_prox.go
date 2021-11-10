package users

import (
	"errors"

	"github.com/neidersalgado/go-camp-grpc/cmd/grpc_server/entities"
)

type UserProx struct{}

func NewUserProxy() *UserProx {
	return &UserProx{}
}

func (up *UserProx) GetUserByID(userId int) (entities.User, error) {

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
