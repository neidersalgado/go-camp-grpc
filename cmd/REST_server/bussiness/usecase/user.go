package usecase

import (
	"context"

	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/entities"
)

type UserService interface {
	Create(ctx context.Context, user entities.User) error
	GetByID(ctx context.Context, userID string) (entities.User, error)
	GetAll(ctx context.Context) ([]entities.User, error)
	Update(ctx context.Context, userToUpdate entities.User) error
	DeleteUser(ctx context.Context, userID string) error
	BulkCreate(ctx context.Context, users *[]entities.User) error
	SetParents(ctx context.Context, userId string, parents *[]entities.User) error
}

type UserUC struct {
	userService UserService
}

func NewUserUseCase(service UserService) *UserUC {
	return &UserUC{
		userService: service,
	}
}

func (uc *UserUC) Create(user entities.User) error {
	return uc.Create(user)
}

func (uc *UserUC) GetByID(userID string) (entities.User, error) {
	return uc.GetByID(userID)
}

func (uc *UserUC) GetAll() (*[]entities.User, error) {
	return uc.GetAll()
}

func (uc *UserUC) Update(userToUpdate entities.User) error {
	return uc.Update(userToUpdate)
}

func (uc *UserUC) DeleteUser(userID string) error {
	return uc.DeleteUser((userID))
}

func (uc *UserUC) BulkCreate(users *[]entities.User) error {
	return uc.BulkCreate(users)
}

func (uc *UserUC) SetParents(userId string, parents *[]entities.User) error {
	return uc.SetParents(userId, parents)
}
