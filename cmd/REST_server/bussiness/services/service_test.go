package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/entities"
	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/repository/mocks"
	service "github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/services"
)

func Test_NewServiceFunction(t *testing.T) {
	userService := service.NewUserService(nil)

	assert.IsType(t, &service.DefaultUserService{}, userService)
}

func Test_CreateUser_WhenCreationIsOk_ThenReturnNilError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	user := getUser()
	repoMock.On("CreateUser", user).Return(nil).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	err := userService.Create(ctx, user)

	assert.Nil(t, err)
	repoMock.AssertExpectations(t)
}

func Test_CreateUser_WhenFails_ThenReturnError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	user := getUser()
	repoError := errors.New("Not Implemented")
	expectedError := fmt.Errorf("Can't create user with ID: %s \n Error: %v", user.Id, repoError)
	repoMock.On("CreateUser", user).Return(repoError).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	err := userService.Create(ctx, user)

	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
	repoMock.AssertExpectations(t)
}

func Test_GetByID_WhenGetOk_ThenReturnUser(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	ExpectedUser := getUser()
	repoMock.On("GetUser", ExpectedUser.Id).Return(ExpectedUser, nil).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	user, err := userService.GetByID(ctx, ExpectedUser.Id)

	assert.Nil(t, err)
	assert.Equal(t, ExpectedUser, user)
	repoMock.AssertExpectations(t)
}

func Test_GetByID_WhenGetNotOk_ThenReturnNilUserAndError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	userID := "1234"
	repoError := errors.New("Not Implemented")
	expectedError := fmt.Errorf("Couldn't Get user with ID: %s \n Error: %v", userID, repoError)
	repoMock.On("GetUser", userID).Return(entities.User{}, repoError).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	user, err := userService.GetByID(ctx, userID)

	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
	assert.Empty(t, user)
	repoMock.AssertExpectations(t)
}

func Test_GetAll_WhenGetIsOk_ThenReturnNilError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	getUsers := getUsers()
	repoMock.On("ListUsers").Return(getUsers, nil).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	users, err := userService.GetAll(ctx)

	assert.Nil(t, err)
	assert.EqualValues(t, users, getUsers)
	assert.Equal(t, getUsers, users)
	repoMock.AssertExpectations(t)
}

func Test_GetAll_WhenGetFails_ThenReturnEmptySliceAndError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	repoError := errors.New("Not Implemented")
	expectedError := fmt.Errorf("Couldn't get users.\n Error: %v", repoError)
	repoMock.On("ListUsers").Return(nil, repoError).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	users, err := userService.GetAll(ctx)

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.EqualError(t, expectedError, err.Error())
	repoMock.AssertExpectations(t)
}

func Test_Update_WhenIsOk_ThenReturnNilError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	user := getUser()
	repoMock.On("UpdateUser", user).Return(nil).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	err := userService.Update(ctx, user)

	assert.Nil(t, err)
	repoMock.AssertExpectations(t)
}

func Test_Update_WhenFails_ThenReturnError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	user := getUser()
	repoError := errors.New("Not Implemented")
	expectedError := fmt.Errorf("Couldn't update user with ID: %s \n Error: %v", user.Id, repoError)
	repoMock.On("UpdateUser", user).Return(repoError).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	err := userService.Update(ctx, user)

	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
	repoMock.AssertExpectations(t)
}

func Test_DeleteUser_WhenDeleteIsOk_ThenReturnNilError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	userID := "1234"
	repoMock.On("DeleteUser", userID).Return(nil).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	err := userService.DeleteUser(ctx, userID)

	assert.Nil(t, err)
	repoMock.AssertExpectations(t)
}

func Test_DeleteUser_WhenDeleteFails_ThenReturnError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	userID := "1234"
	repoError := errors.New("Not Implemented")
	expectedError := fmt.Errorf("Couldn't delete the user with ID: %s \n Error: %v", userID, repoError)
	repoMock.On("DeleteUser", userID).Return(repoError).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	err := userService.DeleteUser(ctx, userID)

	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
	repoMock.AssertExpectations(t)
}

func Test_BulkCreate_WhenBulkIsOk_ThenReturnNilError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	users := getUsers()
	user := getUser()
	repoMock.On("CreateUser", user).Return(nil).Times(3)
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	err := userService.BulkCreate(ctx, &users)

	assert.Nil(t, err)
	repoMock.AssertExpectations(t)
}

func Test_BulkCreate_WhenBulkNotOk_ThenReturnError(t *testing.T) {
	repoMock := mocks.RepositoryMock{}
	users := getUsers()
	user := getUser()
	repoError := errors.New("Not Implemented")
	expectedError := fmt.Errorf("Couldn't create user with ID: %s \n Error: %v", user.Id, repoError)
	repoMock.On("CreateUser", user).Return(repoError).Once()
	userService := service.NewUserService(repoMock)
	ctx := context.TODO()

	err := userService.BulkCreate(ctx, &users)

	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
	repoMock.AssertExpectations(t)
}

func getUser() entities.User {
	return entities.User{
		Id:                    "1234",
		EMail:                 "fake@fake.com",
		Name:                  "fakeName",
		AdditionalInformation: "none",
		Parents:               []entities.User{},
	}
}

func getUsers() []entities.User {
	return []entities.User{
		getUser(),
		getUser(),
		getUser(),
	}
}
