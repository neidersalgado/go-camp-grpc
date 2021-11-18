package main

import "github.com/neidersalgado/go-camp-grpc/cmd/REST_server/transport/models"

type (
	CreateUserRequest struct {
		user models.User
	}
	CreateUserResponse struct {
		Msg string `json:"msg`
		Err error  `json:error, omitempty`
	}
	GetUserRequest struct {
		UserID string `json:user_id`
	}
	GetUserResponse struct {
		User models.User `json:"user,omitempty`
		Err  error       `json:error, omitempty`
	}
	DeleteUserRequest struct {
		UserID string `json:user_id`
	}
	DeleteUserResponse struct {
		Msg string `json:"msg`
		Err error  `json:error, omitempty`
	}
)
