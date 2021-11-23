package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (c *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	//TODO
	// AuthenticateUser(ctx context.Context, email string, hash string) (bool, error)
}

func (c *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	userRequest := UserRequest{}
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userRequest); err != nil {
		responseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	userEntity := transformUserRequestToEntity(userRequest)

	err := c.service.UpdateUser(ctx, userEntity)

	if err != nil {
		responseError(w, http.StatusNotFound, err.Error())
		return
	}

	responseJSON(w, http.StatusNoContent, fmt.Sprintf("Updated User with ID: %v", userRequest.UserID))
}

func (c *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := r.Context()
	id, varsOk := vars["id"]

	if !varsOk {
		responseError(w, http.StatusBadRequest, "Invalid Id")
		return
	}

	userID, err := strconv.Atoi(id)

	if err != nil {
		responseError(w, http.StatusBadRequest, "Invalid id request")
		return
	}

	user, err := c.service.GetUser(ctx, int32(userID))

	if err != nil {
		responseError(w, http.StatusNotFound, err.Error())
	}

	responseJSON(w, http.StatusOK, user)
}

func (c *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := c.service.GetAllUsers(ctx)

	if err != nil {
		responseError(w, http.StatusNotFound, err.Error())
	}

	responseJSON(w, http.StatusOK, users)
}

func (c *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		responseError(w, http.StatusBadRequest, "Invalid body request")
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		responseError(w, http.StatusBadRequest, "Invalid id request")
		return
	}
	err = c.service.DeleteUser(ctx, int32(userID))

	if err != nil {
		responseError(w, http.StatusNotFound, err.Error())
		return
	}

	responseJSON(w, http.StatusNoContent, fmt.Sprintf("deleted User with ID: %v", id))
}
func (c *UserHandler) BulkCreateUser(w http.ResponseWriter, r *http.Request) {
	//	BulkCreateUser(ctx context.Context, users []entities.User) error
}
func (c *UserHandler) SetUserParents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := r.Context()
	id, varsOk := vars["id"]
	userRequest := UserRequest{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userRequest); err != nil {
		responseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if !varsOk {
		responseError(w, http.StatusBadRequest, "Invalid Id")
		return
	}

	userID, err := strconv.Atoi(id)

	if err != nil {
		responseError(w, http.StatusBadRequest, "Invalid id request")
		return
	}

	defer r.Body.Close()
	userEntity := transformUserRequestToEntity(userRequest)

	err = c.service.SetUserParents(ctx, int32(userID), userEntity.Parent)

	if err != nil {
		responseError(w, http.StatusNotFound, err.Error())
	}

	responseJSON(w, http.StatusOK, fmt.Sprintf("set parents to User with ID: %v", id))
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
