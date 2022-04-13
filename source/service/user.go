package service

import (
	"app/source/controllers/requests"
	"app/source/controllers/responses"
	"app/source/domain/entities"
	"app/source/domain/exception"
	"app/source/repository"
	"app/source/utils"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
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
}

func InsertUser(userRequest *requests.UserRequest) {

	prepareUser(userRequest)

	user := entities.User{
		Login:     userRequest.Login,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		CreatedAt: time.Now(),
		UpdateAt:  time.Time{},
	}

	err := repository.InsertUser(&user)

	if err != nil {
		log.Panic("<InsertPlanet> An error ocurred during insert", err)
	}
}

func FindUserById(id int) *responses.UserResponse {
	result, err := repository.FindUserById(id)

	if err != nil {
		panic(exception.NewNotFoundException(fmt.Sprintf("User id %d was not found", id)))
	}

	userResponse := mapUserEntityToResponse(result)
	return &userResponse
}

func UpdateUser(userRequest *requests.UserRequest, id int) {

	prepareUser(userRequest)
	userFounded := FindUserById(id)

	user := entities.User{
		ID:       userFounded.ID,
		Login:    userRequest.Login,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	err := repository.UpdateUser(&user)

	if err != nil {
		log.Panic("<UpdateUser> An error ocurred during update", err)
	}
}

func DeleteUser(id int) {
	err := repository.DeleteUser(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(&exception.HttpException{StatusCode: http.StatusNotFound,
				Message: fmt.Sprintf("Planet cannot deleted because id %d was not found", id)})
		}
		log.Panic("An error ocurred during delete", err)
	}
}

func mapUserEntityToResponse(planet *entities.User) (response responses.UserResponse) {

	return responses.UserResponse{
		ID:        response.ID,
		Login:     response.Login,
		Email:     response.Email,
		Password:  response.Password,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
}
