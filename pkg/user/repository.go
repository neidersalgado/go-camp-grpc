package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/neidersalgado/go-camp-grpc/pkg/entities"
	"github.com/neidersalgado/go-camp-grpc/pkg/user/pb"
)

type ProxyRepository struct {
}

func NewProxyRepository() *ProxyRepository {
	return &ProxyRepository{}
}

func (up ProxyRepository) Authenticate(email string, hash string) (bool, error) {
	fmt.Println("User well authenticated")
	return true, nil
}

func (up ProxyRepository) Create(user entities.User) error {
	serverCon, err := OpenServerConnection()

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
	}

	defer serverCon.dispose()
	c := serverCon.client
	externalUser := transformUserEntityToRequest(user)
	result, errorFromCall := c.Create(serverCon.context, &externalUser)
	if errorFromCall != nil {
		fmt.Sprint(fmt.Sprintf("\n Error Creating User  error: %v \n", errorFromCall.Error()))
		return errors.New(fmt.Sprintf("Error Creating User  error: %v", errorFromCall))
	}

	if result.GetCode() != http.StatusOK {
		return errors.New(fmt.Sprintf("Error Creating User  error: %v, response code: %v", errorFromCall, result.Code))
	}

	return nil
}

func (up ProxyRepository) Update(user entities.User) error {
	fmt.Println("User well Updated, mock")
	fmt.Println(fmt.Sprintf("user: %v", user))
	return nil
}

func (up ProxyRepository) Get(ctx context.Context, userID int32) (entities.User, error) {

	serverCon, err := OpenServerConnection()

	if err != nil {
		log.Fatalf("did not connect to server: %s", err)
	}

	defer serverCon.dispose()
	c := serverCon.client
	userIDpb := transformUserIdToUserIdRequest(userID)
	userResponse, errorFromCall := c.Get(ctx, userIDpb)

	if errorFromCall != nil {
		fmt.Sprint(fmt.Sprintf("\n Error Creating User  error: %v \n", errorFromCall.Error()))
		return entities.User{}, errors.New(fmt.Sprintf("Error Creating User  error: %v", errorFromCall.Error()))
	}

	userEntity := transformUserResponseToEntity(*userResponse)

	return userEntity, nil
}

func (up ProxyRepository) List() ([]entities.User, error) {
	return []entities.User{}, nil
	//[]entities.User{}, fmt.Errorf("List Users Not Implemented")
}

func (up ProxyRepository) Delete(userID int32) error {
	return nil
	//fmt.Errorf("Delete Users Not Implemented")
}

func OpenServerConnection() (*ServerConnection, error) {

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
