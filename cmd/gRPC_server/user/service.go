package users

import (
	"github.com/neidersalgado/go-camp-grpc/cmd/gRPC_server/pb"
)

type Repository interface {
	Create(user pb.UserRequest) error
	Get(userID pb.UserID) (pb.UserResponse, error)
	Update(user pb.UserRequest) error
	Delete(userID pb.UserID) error
	GetAll() (pb.UserColletionResponse, error)
}

type UsersService struct {
	pb.UnimplementedUsersServer
	repository Repository
}

func NewUserService(repo Repository) *UsersService {
	return &UsersService{
		repository: repo,
	}
}
