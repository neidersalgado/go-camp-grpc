package usecase

import "github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/entities"

type userService interface {
	Create(user entities.User) error
	GetByID(userID string) (entities.User, error)
	GetAll() ([]entities.User, error)
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
