package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	users "github.com/neidersalgado/go-camp-grpc/cmd/grpc_server/application"
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

func responserError(w http.ResponseWriter, code int, message string) {
	responseJSON(w, code, map[string]string{"error": message})
}

func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
