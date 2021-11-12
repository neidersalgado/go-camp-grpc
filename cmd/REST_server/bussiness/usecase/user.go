package usecase

import "github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/entities"

type userService interface {
	Create(user entities.User) error
	GetByID(userID string) (entities.User, error)
	GetAll() (*[]entities.User, error)
	Update(userToUpdate entities.User) error
	DeleteUser(userID string) error
	BulkCreate(users *[]entities.User) error
	SetParents(userId string, parents *[]entities.User) error
}

type UserUC struct {
	userService userService
}

func NewUserUseCase(service userService) *UserUC {
	return &UserUC{
		userService: service,
	}
}

