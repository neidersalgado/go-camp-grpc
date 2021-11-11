package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	users "github.com/neidersalgado/go-camp-grpc/cmd/grpc_server/application"
	"github.com/neidersalgado/go-camp-grpc/cmd/grpc_server/entities"
)

type UserController struct {
	dataSource *users.UserProx
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, varsOk := vars["id"]

	userId, err := strconv.Atoi(id)

	if !varsOk || err != nil {
		responserError(w, http.StatusBadRequest, "Invalid Id")
		return
	}

	user, err := c.dataSource.GetUserByID(userId)

	if err != nil {
		responserError(w, http.StatusNotFound, err.Error())
	}

	responseJSON(w, http.StatusOK, user)

}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := c.dataSource.GetAll()

	if err != nil {
		responserError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := resp

	responseJSON(w, http.StatusOK, response)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {

	userToCreate := entities.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userToCreate); err != nil {
		responserError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	userCreated, err := c.dataSource.Create(userToCreate)

	if err != nil {
		responserError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusCreated, userCreated)

}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {

	userToUpdate := entities.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userToUpdate); err != nil {
		responserError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	vars := mux.Vars(r)

	email, ok := vars["id"]

	if !ok {
		responserError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	userToUpdate.EMail = email

	defer r.Body.Close()

	if err := c.dataSource.Update(userToUpdate); err != nil {
		responserError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusNoContent, userToUpdate)

}
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		responserError(w, http.StatusBadRequest, "Invalid body request")
		return
	}

	intId, res := strconv.Atoi(id)

	if res != nil {
		responserError(w, http.StatusBadRequest, "Invalid  is format")
		return
	}

	err := c.dataSource.Delete(int(intId))

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
