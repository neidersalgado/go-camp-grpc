package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/models"
	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/service"
)

func main() {
	r := mux.NewRouter()
	proxyRepoService := service.NewUserProxy()
	userSvc := service.NewUserService(proxyRepoService)

	CreateUserHandler := httptransport.NewServer(
		makeCreateUserEndpoint(userSvc),
		decodeCreateUserRequest,
		encodeResponse,
	)
	http.Handle("/", r)
	http.Handle("/users", CreateUserHandler)
	r.Methods("GET").Path("/users").Handler(CreateUserHandler)
	fmt.Println("msg", "HTTP", "addr", "8081")
	fmt.Println("err", http.ListenAndServe(":8081", r))
}

func makeCreateUserEndpoint(s UserService) endpoint.Endpoint {
	fmt.Println("make createUser endpoint called")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		err := s.Create(ctx, req.User)
		return CreateUserResponse{Msg: req.User.Id, Err: err}, nil
	}
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("decodeCreateUserRequest called")
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req.User); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("Encode Response called")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type UserService interface {
	Create(ctx context.Context, user models.User) error
}

type (
	CreateUserRequest struct {
		User models.User
	}
	CreateUserResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error", omitempty`
	}
	GetUserRequest struct {
		UserID string `json:"user_id"`
	}
	GetUserResponse struct {
		User models.User `json:"user",omitempty`
		Err  error       `json:"error", omitempty`
	}
	DeleteUserRequest struct {
		UserID string `json:"user_id"`
	}
	DeleteUserResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error", omitempty`
	}
)
