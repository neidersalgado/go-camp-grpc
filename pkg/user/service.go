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
	GetUser(ctx context.Context, userID string) (entities.User, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
	DeleteUser(ctx context.Context, userID string) error
	BulkCreateUser(ctx context.Context, users []entities.User) error
	SetUserParents(ctx context.Context, userID string, parents []entities.User) error
}

type Repository interface {
	Authenticate(email string, hash string) (bool, error)
	Create(user entities.User) error
	Update(user entities.User) error
	Get(ctx context.Context, userID string) (entities.User, error)
	List(ctx context.Context) ([]entities.User, error)
	Delete(ctx context.Context, userID string) error
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
	//logger := log.With(s.logger, "method", "Create")

	if err := s.repository.Create(user); err != nil {
		//	level.Error(logger).Log("Error from repo is %v", err.Error())
		fmt.Sprintf("Error Creating User service error: %v", err.Error())
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

func (s DefaultUserService) GetUser(ctx context.Context, userID string) (entities.User, error) {
	//logger := log.With(s.logger, "method", "GetUserByID")
	fmt.Printf("service.GerUser with id: %v.\n", userID)
	user, err := s.repository.Get(ctx, userID)

	if err != nil {
		//level.Error(logger).Log("err from repo is", err.Error())
		fmt.Sprintf("Error Get User  Service error: %v", err.Error())
		return entities.User{}, err
	}

	return user, nil
}

func (s DefaultUserService) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	//logger := log.With(s.logger, "method", "GetAllUsers")
	fmt.Printf("service.getAll with id: \n")
	users, err := s.repository.List(ctx)

	if err != nil {
		fmt.Print(err)
		//level.Error(logger).Log("err from repo is", err)
		fmt.Sprintf("Error Get Users  Service error: %v", err.Error())
		return users, err
	}

	return users, nil
}

func (s DefaultUserService) DeleteUser(ctx context.Context, userID string) error {
	logger := log.With(s.logger, "method", "DeleteUser")
	err := s.repository.Delete(ctx, userID)

	if err != nil {
		level.Error(logger).Log("err from repo is", err)
		return err
	}

	return nil
}

func (s DefaultUserService) BulkCreateUser(ctx context.Context, users []entities.User) error {
	//logger := log.With(s.logger, "method", "BulkUsers")

	for _, user := range users {
		err := s.repository.Create(user)

		if err != nil {
			//level.Error(logger).Log("err from repo is", err)
			return err
		}
	}

	return nil
}
func (s DefaultUserService) SetUserParents(ctx context.Context, userID string, parents []entities.User) error {
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
