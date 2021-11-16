package endpoint

import (
	"github.com/gorilla/mux"

	service "github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/services"
	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/usecase"
	infrastructure "github.com/neidersalgado/go-camp-grpc/cmd/REST_server/transport/repository"
)

func SetUpRouter(router *mux.Router) {
	repository := infrastructure.NewUserRepository()
	userService := service.NewUserService(repository)
	useCase := usecase.NewUserUseCase(userService)
	usersController := NewUserController(*useCase)
	router.HandleFunc("/users/{id}", usersController.GetByID).Methods("GET")
	/**
	router.HandleFunc("/users", usersController.GetAll).Methods("GET")
	router.HandleFunc("/users", usersController.Create).Methods("POST")
	router.HandleFunc("/users/{id}", usersController.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", usersController.Delete).Methods("DELETE")
	**/
}
