package endpoint

import (
	"github.com/gorilla/mux"
)

func SetUpRouter(router *mux.Router) {
	usersController := NewUserController()
	router.HandleFunc("/users/{id}", usersController.GetByID).Methods("GET")
	router.HandleFunc("/users", usersController.GetAll).Methods("GET")
	router.HandleFunc("/users", usersController.Create).Methods("POST")
	router.HandleFunc("/users/{id}", usersController.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", usersController.Delete).Methods("DELETE")
}
