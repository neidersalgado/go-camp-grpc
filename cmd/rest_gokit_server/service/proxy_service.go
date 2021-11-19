package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/models"
	service "github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/service/pb"
)

type UserProxy struct {
}

func NewUserProxy() *UserProxy {
	return &UserProxy{}
}

func (up UserProxy) CreateUser(u models.User) error {

	serverCon, err := OpenServerConnection()

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
	}

	defer serverCon.dispose()
	c := serverCon.client
	externalUser := &service.UserRequest{
		Id:                    u.Id,
		Name:                  u.Name,
		PwdHash:               u.PwdHash,
		AdditionalInformation: u.AdditionalInformation,
		Age:                   u.Age,
	}

	result, errorFromCall := c.Create(serverCon.context, externalUser)

	if result.Code != service.Response_OK {
		return errors.New("Error Creating User")
	}

	return errorFromCall
}

func (up UserProxy) DeleteUser(userID string) error {

	serverCon, err := OpenServerConnection()

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
	}

	defer serverCon.dispose()
	c := serverCon.client
	externalUserId := &service.UserID{
		ID: userID,
	}
	result, errorFromCall := c.Delete(serverCon.context, externalUserId)

	if result.Code == service.Response_OK {
		return nil
	}

	return errorFromCall
}

func (up UserProxy) GetUser(userID string) (models.User, error) {

	serverCon, err := OpenServerConnection()

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
	}

	defer serverCon.dispose()
	c := serverCon.client
	externalUserId := &service.UserID{
		ID: userID,
	}
	userFromGrpc, errorFromCall := c.Get(serverCon.context, externalUserId)

	if errorFromCall != nil {
		fmt.Println("server call did not work:", errorFromCall)
		return models.User{}, errorFromCall
	}

	response := models.User{
		Id:                    userFromGrpc.Id,
		Name:                  userFromGrpc.Name,
		PwdHash:               userFromGrpc.PwdHash,
		AdditionalInformation: userFromGrpc.AdditionalInformation,
	}

	return response, errorFromCall
}

func (up UserProxy) ListUsers() ([]models.User, error) {
	return []models.User{}, fmt.Errorf("List Users Not Implemented")
}
func (up UserProxy) UpdateUser(user models.User) error {
	return fmt.Errorf("UpdateUser Not Implemented")
}

func OpenServerConnection() (*ServerConnection, error) {

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
		return nil, err //unreached?
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	c := service.NewUsersClient(conn)

	return &ServerConnection{client: c, context: ctx, dispose: func() {
		cancel()
		conn.Close()

	}}, nil

}

type ServerConnection struct {
	client  service.UsersClient
	context context.Context
	dispose func()
}
