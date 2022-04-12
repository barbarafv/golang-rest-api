package service

import (
	"app/source/controllers/requests"
	"app/source/controllers/responses"
	"app/source/domain/entities"
	"app/source/repository"
	"app/source/utils"
	"html"
	"log"
	"strings"
	"time"
)

func BeforeSaveUser(request *requests.UserRequest) error {
	hashedPassword, err := utils.Hash(request.Password)
	if err != nil {
		return err
	}
	request.Password = string(hashedPassword)
	return nil
}

func prepareUser(request *requests.UserRequest) {
	request.Email = html.EscapeString(strings.TrimSpace(request.Email))
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
}

func InsertUser(userRequest *requests.UserRequest) {

	prepareUser(userRequest)

	user := mapUserRequestToEntity(userRequest)

	err := repository.InsertUser(&user)

	if err != nil {
		log.Panic("<InsertPlanet> An error ocurred during insert", err)
	}
}

func mapUserEntityToResponse(id string, planet *entities.User) (response responses.UserResponse) {

	return responses.UserResponse{
		Id:        response.Id,
		Login:     response.Login,
		Email:     response.Email,
		Password:  response.Password,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}

}

func mapUserRequestToEntity(request *requests.UserRequest) entities.User {

	return entities.User{
		Login:     request.Login,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: request.CreatedAt,
		UpdateAt:  request.UpdatedAt,
	}
}
