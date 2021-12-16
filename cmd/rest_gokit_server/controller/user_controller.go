package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/models"
	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/service"
)

type UserController struct {
	dataSource *service.UserProxy
}

func NewUserController(service *service.UserProxy) *UserController {
	return &UserController{dataSource: service}
}
func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, varsOk := vars["id"]
	if !varsOk {
		responserError(w, http.StatusBadRequest, "Invalid Id")
		return
	}
	user, err := c.dataSource.GetUser(userID)

	if err != nil {
		responserError(w, http.StatusNotFound, err.Error())
	}
	responseJSON(w, http.StatusOK, user)

}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {

	userToCreate := models.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userToCreate); err != nil {
		responserError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	err := c.dataSource.CreateUser(userToCreate)

	if err != nil {
		responserError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusCreated, userToCreate)

}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		responserError(w, http.StatusBadRequest, "Invalid body request")
		return
	}

	err := c.dataSource.DeleteUser(id)

	if err != nil {
		responserError(w, http.StatusNotFound, err.Error())
		return
	}

	responseJSON(w, http.StatusNoContent, fmt.Sprintf("Unable to delete User with ID: %v", id))
}

//TODo edit to response no errors just messages
func responserError(w http.ResponseWriter, code int, message string) {
	responseJSON(w, code, map[string]string{"error": message})
}

//create function response with status
func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
