package main

import (
	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/entities"
	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/transport/models"
)

func EntityToModel(user entities.User) models.User {
	return models.User{
		Id:                    user.Id,
		Name:                  user.Name,
		PwdHash:               user.PwdHash,
		Age:                   user.Age,
		AdditionalInformation: user.AdditionalInformation,
	}
}

func ModelToEntity(user models.User) entities.User {
	return entities.User{
		Id:                    user.Id,
		Name:                  user.Name,
		PwdHash:               user.PwdHash,
		Age:                   user.Age,
		AdditionalInformation: user.AdditionalInformation,
	}
}
