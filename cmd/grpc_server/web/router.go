package web

import (
	"github.com/gorilla/mux"
)

func SetUpRouter(router *mux.Router) {
	usersController := NewUserController()
	router.HandleFunc("/users/{id}", usersController.GetByID).Methods("GET")
}
