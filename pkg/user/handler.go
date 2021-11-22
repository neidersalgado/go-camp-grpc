package user

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (c *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	userRequest := UserRequest{}
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userRequest); err != nil {
		responseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	userEntity := transformUserRequestToEntity(userRequest)
	err := c.service.CreateUser(ctx, userEntity)

	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusCreated, userRequest)
}

//TODo edit to response no errors just messages
func responseError(w http.ResponseWriter, code int, message string) {
	responseJSON(w, code, map[string]string{"error": message})
}

//create function response with status
func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
