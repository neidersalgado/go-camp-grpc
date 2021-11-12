package endpoint

import (
	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/usecase"
)

type UserController struct {
	uc usecase.UserUC
}

func NewUserController(usecase usecase.UserUC) *UserController {
	return &UserController{
		uc: usecase,
	}
}

/*
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

	responseMessage(w, http.StatusCreated, fmt.Sprint("User With Id:%D Created.", userCreated.Id))

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

	responseMessage(w, http.StatusNoContent, fmt.Sprintf("User With Id: %d updated correctly.", userToUpdate.Id))

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
		responserError(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err := c.dataSource.Delete(int(intId))

	if err != nil {
		responserError(w, http.StatusNotFound, err.Error())
		return
	}

	responseMessage(w, http.StatusNoContent, fmt.Sprintf("Deleted user with ID: %v", id))
}

func responserError(w http.ResponseWriter, code int, message string) {
	responseJSON(w, code, map[string]string{"error": message})
}

func responseMessage(w http.ResponseWriter, code int, message string) {
	responseJSON(w, code, map[string]string{"response": message})
}

func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
*/
