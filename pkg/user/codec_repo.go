package user

import "github.com/neidersalgado/go-camp-grpc/pkg/entities"

func transformUserEntityToModel(userEntity entities.User) UserModel {
	return UserModel{
		UserID:                userEntity.UserID,
		Email:                 userEntity.Email,
		Name:                  userEntity.Name,
		Age:                   userEntity.Age,
		AdditionalInformation: userEntity.AdditionalInformation,
		Parent:                transformUserEntitiesToModels(userEntity.Parent),
	}
}

func transformUserEntitiesToModels(userEntities []entities.User) []UserModel {
	if len(userEntities) == 0 {
		return []UserModel{}
	}

	parentsModel := make([]UserModel, len(userEntities))

	for _, entity := range userEntities {
		parent := transformUserEntityToModel(entity)
		parentsModel = append(parentsModel, parent)
	}

	return parentsModel
}

func transformUserModelToEntity(userModel UserModel) entities.User {
	return entities.User{
		UserID:                userModel.UserID,
		Email:                 userModel.Email,
		Name:                  userModel.Name,
		Age:                   userModel.Age,
		AdditionalInformation: userModel.AdditionalInformation,
		Parent:                transformUserModelsToEntities(userModel.Parent),
	}
}

func transformUserModelsToEntities(userModels []UserModel) []entities.User {
	if len(userModels) == 0 {
		return []entities.User{}
	}

	parentsEntities := make([]entities.User, len(userModels))

	for _, model := range userModels {
		parent := transformUserModelToEntity(model)
		parentsEntities = append(parentsEntities, parent)
	}

	return parentsEntities
}
