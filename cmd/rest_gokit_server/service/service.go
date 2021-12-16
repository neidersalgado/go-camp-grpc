package service

import (
	"context"
	"fmt"

	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/models"
	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/repository"
)

type DefaultUserService struct {
	repository repository.Repository
}

func NewUserService(repo repository.Repository) *DefaultUserService {
	return &DefaultUserService{
		repository: repo,
	}
}

func (s *DefaultUserService) Create(ctx context.Context, user models.User) error {
	if err := s.repository.CreateUser(user); err != nil {
		return fmt.Errorf("Can't create user with ID: %s \n Error: %v", user.Id, err)
	}

	return nil
}

func (s *DefaultUserService) GetByID(ctx context.Context, userID string) (models.User, error) {
	user, err := s.repository.GetUser(userID)

	if err != nil {
		return models.User{}, fmt.Errorf("Couldn't Get user with ID: %s \n Error: %v", userID, err)
	}

	return user, nil
}

func (s *DefaultUserService) GetAll(ctx context.Context) ([]models.User, error) {
	users, err := s.repository.ListUsers()

	if err != nil {
		return nil, fmt.Errorf("Couldn't get users.\n Error: %v", err)
	}

	return users, nil
}

func (s *DefaultUserService) Update(ctx context.Context, userToUpdate models.User) error {
	if err := s.repository.UpdateUser(userToUpdate); err != nil {
		return fmt.Errorf("Couldn't update user with ID: %s \n Error: %v", userToUpdate.Id, err)
	}

	return nil
}

func (s *DefaultUserService) DeleteUser(ctx context.Context, userID string) error {
	if err := s.repository.DeleteUser(userID); err != nil {
		return fmt.Errorf("Couldn't delete the user with ID: %s \n Error: %v", userID, err)
	}

	return nil
}

func (s *DefaultUserService) BulkCreate(ctx context.Context, users *[]models.User) error {
	for _, user := range *users {
		if err := s.repository.CreateUser(user); err != nil {
			return fmt.Errorf("Couldn't create user with ID: %s \n Error: %v", user.Id, err)
		}
	}

	return nil
}

func (s *DefaultUserService) SetParents(ctx context.Context, userId string, parents *[]models.User) error {
	user, err := s.GetByID(ctx, userId)

	if err != nil {
		return fmt.Errorf("Couldn't get user with ID: %s  to update parents.\n Error: %v", userId, err)
	}

	user.Parents = append(user.Parents, *parents...)

	if err := s.repository.UpdateUser(user); err != nil {
		return fmt.Errorf("Couldn't update parents for user with ID: %s \n Error: %v", userId, err)
	}

	return nil
}
