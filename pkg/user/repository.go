package user

import (
	"fmt"

	"github.com/neidersalgado/go-camp-grpc/pkg/entities"
)

type ProxyRepository struct {
}

func NewProxyRepository() *ProxyRepository {
	return &ProxyRepository{}
}

func (up ProxyRepository) Authenticate(email string, hash string) (bool, error) {
	fmt.Println("User well authenticated")
	return true, nil
}

func (up ProxyRepository) Create(user entities.User) error {
	fmt.Println("User well Create")
	userModel := transformUserEntityToModel(user)
	fmt.Println(fmt.Sprintf("user: %v", userModel))
	return nil
}

func (up ProxyRepository) Update(user entities.User) error {
	fmt.Println("User well Updated")
	fmt.Println(fmt.Sprintf("user: %v", user))
	return nil
}

func (up ProxyRepository) Get(userID int32) (entities.User, error) {
	fmt.Println("User well Get")
	userModel := UserModel{
		UserID:                userID,
		PWDHash:               "H4$h",
		Email:                 "mail@mail.com",
		Name:                  "name lasname",
		Age:                   "23",
		AdditionalInformation: "none",
		Parent:                nil,
	}
	// TODO cast from model to entity
	userEntity := transformUserModelToEntity(userModel)
	return userEntity, nil
}

func (up ProxyRepository) List() ([]entities.User, error) {
	return []entities.User{}, fmt.Errorf("List Users Not Implemented")
}

func (up ProxyRepository) Delete(userID int32) error {
	return fmt.Errorf("Delete Users Not Implemented")
}
