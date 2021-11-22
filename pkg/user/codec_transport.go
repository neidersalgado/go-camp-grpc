package user

import "github.com/neidersalgado/go-camp-grpc/pkg/entities"

func transformUserEntityToRequest(userEntity entities.User) UserRequest {
	return UserRequest{
		UserID:                userEntity.UserID,
		Email:                 userEntity.Email,
		Name:                  userEntity.Name,
		Age:                   userEntity.Age,
		AdditionalInformation: userEntity.AdditionalInformation,
		Parent:                transformUserEntitiesToRequests(userEntity.Parent),
	}
}

func transformUserEntitiesToRequests(userEntities []entities.User) []UserRequest {
	if len(userEntities) == 0 {
		return []UserRequest{}
	}

	parentsRequest := make([]UserRequest, len(userEntities))

	for _, entity := range userEntities {
		parent := transformUserEntityToRequest(entity)
		parentsRequest = append(parentsRequest, parent)
	}

	return parentsRequest
}

func transformUserRequestToEntity(userRequest UserRequest) entities.User {
	return entities.User{
		UserID:                userRequest.UserID,
		Email:                 userRequest.Email,
		Name:                  userRequest.Name,
		Age:                   userRequest.Age,
		AdditionalInformation: userRequest.AdditionalInformation,
		Parent:                transformUserRequestsToEntities(userRequest.Parent),
	}
}

func transformUserRequestsToEntities(userRequest []UserRequest) []entities.User {
	if len(userRequest) == 0 {
		return []entities.User{}
	}

	parentsEntities := make([]entities.User, len(userRequest))

	for _, entity := range userRequest {
		parent := transformUserRequestToEntity(entity)
		parentsEntities = append(parentsEntities, parent)
	}

	return parentsEntities
}
