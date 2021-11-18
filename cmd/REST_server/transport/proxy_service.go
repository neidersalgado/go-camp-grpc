package transport

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/entities"
	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/pb"
)

type UserProxy struct {
}

func NewUserProxy() *UserProxy {
	return &UserProxy{}
}

func (up UserProxy) Create(u entities.User) (entities.User, error) {

	serverCon, err := OpenServerConection()

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
	}

	defer serverCon.dispose()
	c := serverCon.client
	externalUser := &pb.UserRequest{
		Id:                    "",
		Name:                  u.Name,
		PwdHash:               u.PwdHash,
		AdditionalInformation: u.AdditionalInformation,
		Age:                   u.Age,
	}

	result, errorFromCall := c.Create(serverCon.context, externalUser)

	if result.Code != pb.Response_OK {
		return entities.User{}, errors.New("error al crear")
	}

	return u, errorFromCall
}

func (up UserProxy) Delete(id string) (bool, error) {

	serverCon, err := OpenServerConection()

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
	}

	defer serverCon.dispose()
	c := serverCon.client
	externalUserId := &pb.UserID{
		ID: id,
	}
	result, errorFromCall := c.Delete(serverCon.context, externalUserId)

	if result.Code == pb.Response_OK {
		return true, nil
	}

	return false, errorFromCall
}

func (up UserProxy) GetById(userID string) (entities.User, error) {

	serverCon, err := OpenServerConection()

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
	}

	defer serverCon.dispose()
	c := serverCon.client
	externalUserId := &pb.UserID{
		ID: userID,
	}
	userFromGrpc, errorFromCall := c.Get(serverCon.context, externalUserId)

	if errorFromCall != nil {
		fmt.Println("server call did not work:", errorFromCall)
		return entities.User{}, errorFromCall
	}

	response := entities.User{
		Id:                    userFromGrpc.Id,
		Name:                  userFromGrpc.Name,
		PwdHash:               userFromGrpc.PwdHash,
		AdditionalInformation: userFromGrpc.AdditionalInformation,
	}

	return response, errorFromCall
}

func OpenServerConection() (*ServerConnection, error) {

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
		return nil, err //unreached?
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	c := pb.NewUsersClient(conn)

	return &ServerConnection{client: c, context: ctx, dispose: func() {
		cancel()
		conn.Close()

	}}, nil

}

type ServerConnection struct {
	client  pb.UsersClient
	context context.Context
	dispose func()
}
