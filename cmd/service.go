package services

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/neidersalgado/go-camp-grpc/cmd/models"
)

type UserService interface {
	CreateUser(ctx context.Context, customer models.User)
	GetUserById(ctx context.Context, id string) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, customer models.User) (string, error)
	DeleteUser(ctx context.Context, id string) (string, error)
}

type Repository interface {
	CreateUser(ctx context.Context, customer models.User) error
	GetUserById(ctx context.Context, id string) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, customer models.User) (string, error)
	DeleteUser(ctx context.Context, id string) (string, error)
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

func (s DefaultUserService) CreateUser(ctx context.Context, user models.User) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	var msg = "success"
	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err from repo is", err)
		return "", err
	}
	return msg, nil
}

func (s DefaultUserService) GetUserById(ctx context.Context, userID string) (models.User, error) {
	logger := log.With(s.logger, "method", "GetUserByID")
	user, err := s.repository.GetUserById(ctx, userID)
	if err != nil {
		level.Error(logger).Log("err from repo is", err)
		return models.User{}, err
	}
	return user, nil
}

func (s DefaultUserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	logger := log.With(s.logger, "method", "GetAllUsers")
	users, err := s.repository.GetAllUsers(ctx)
	if err != nil {
		level.Error(logger).Log("err from repo is", err)
		return []models.User{}, err
	}
	return users, nil
}

func (s DefaultUserService) UpdateUser(ctx context.Context, user models.User) (string, error) {
	logger := log.With(s.logger, "method", "UpdateUSer")
	msg, err := s.repository.UpdateUser(ctx, user)
	if err != nil {
		level.Error(logger).Log("err from repo is", err)
		return "", err
	}
	return msg, nil
}

func (s DefaultUserService) DeleteUser(ctx context.Context, userID string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteUser")
	msg, err := s.repository.DeleteUser(ctx, userID)
	if err != nil {
		level.Error(logger).Log("err from repo is", err)
		return "", err
	}
	return msg, nil
}
