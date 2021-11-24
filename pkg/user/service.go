package user

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/neidersalgado/go-camp-grpc/pkg/entities"
)

type UserService interface {
	AuthenticateUser(ctx context.Context, email string, hash string) (bool, error)
	CreateUser(ctx context.Context, user entities.User) error
	UpdateUser(ctx context.Context, user entities.User) error
	GetUser(ctx context.Context, userID int32) (entities.User, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
	DeleteUser(ctx context.Context, userID int32) error
	BulkCreateUser(ctx context.Context, users []entities.User) error
	SetUserParents(ctx context.Context, userID int32, parents []entities.User) error
}

type Repository interface {
	Authenticate(email string, hash string) (bool, error)
	Create(user entities.User) error
	Update(user entities.User) error
	Get(ctx context.Context, userID int32) (entities.User, error)
	List() ([]entities.User, error)
	Delete(userID int32) error
}

type DefaultUserService struct {
	repository Repository
	logger     log.Logger
}

func NewDefaultUserService(repository Repository) *DefaultUserService {
	return &DefaultUserService{
		repository: repository,
	}
}

func (s DefaultUserService) AuthenticateUser(ctx context.Context, email string, hash string) (bool, error) {
	return false, nil
}

func (s DefaultUserService) CreateUser(ctx context.Context, user entities.User) error {
	logger := log.With(s.logger, "method", "Create")

	if err := s.repository.Create(user); err != nil {
		fmt.Println("*************ERROR ************")
		fmt.Println("Error occur wile creating %v", err.Error())
		level.Error(logger).Log("Error from repo is %v", err.Error())
		return err
	}

	return nil
}

func (s DefaultUserService) UpdateUser(ctx context.Context, user entities.User) error {
	logger := log.With(s.logger, "method", "UpdateUSer")
	err := s.repository.Update(user)

	if err != nil {
		level.Error(logger).Log("err from repo is", err)
		return err
	}

	return nil
}

func (s DefaultUserService) GetUser(ctx context.Context, userID int32) (entities.User, error) {
	logger := log.With(s.logger, "method", "GetUserByID")
	user, err := s.repository.Get(ctx, userID)

	if err != nil {
		level.Error(logger).Log("err from repo is", err.Error())
		return entities.User{}, err
	}

	return user, nil
}

func (s DefaultUserService) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	logger := log.With(s.logger, "method", "GetAllUsers")
	users, err := s.repository.List()

	if err != nil {
		fmt.Print(err)
		level.Error(logger).Log("err from repo is", err)
		return users, err
	}

	return users, nil
}

func (s DefaultUserService) DeleteUser(ctx context.Context, userID int32) error {
	logger := log.With(s.logger, "method", "DeleteUser")
	err := s.repository.Delete(userID)

	if err != nil {
		level.Error(logger).Log("err from repo is", err)
		return err
	}

	return nil
}

func (s DefaultUserService) BulkCreateUser(ctx context.Context, users []entities.User) error {
	logger := log.With(s.logger, "method", "BulkUsers")

	for _, user := range users {
		err := s.repository.Create(user)

		if err != nil {
			level.Error(logger).Log("err from repo is", err)
			return err
		}
	}

	return nil
}
func (s DefaultUserService) SetUserParents(ctx context.Context, userID int32, parents []entities.User) error {
	logger := log.With(s.logger, "method", "SetUserParents")
	user, err := s.repository.Get(ctx, userID)

	if err != nil {
		level.Error(logger).Log("err from repo is", err)
		return err
	}

	user.Parent = append(user.Parent, parents...)
	err = s.repository.Update(user)

	if err != nil {
		level.Error(logger).Log("err from repo is", err)
		return err
	}

	return nil
}
