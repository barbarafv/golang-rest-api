package tests

import (
	_ "app/source/_testinit/fixture"
	"app/source/controllers/requests"
	"app/source/domain/entities"
	"app/source/repository"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.POST("/users").SetJSONInterface(requests.UserRequest{
			Login:    "login",
			Email:    "login@email.com",
			Password: "login1234",
		}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, response.Code)

			user, _ := repository.FindUserById(1)

			assert.Equal(t, "login", user.Login)
			assert.Equal(t, "login@email.com", user.Email)
			assert.Equal(t, "login1234", user.Password)

		})
	})
}

func TestGetUserById(t *testing.T) {
	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		repository.InsertUser(&entities.User{ID: 1, Login: "user123", Email: "login123@email.com",
			Password: "123456"})

		rest.GET("/users/1").SetJSONInterface(requests.UserRequest{
			Login:    "login",
			Email:    "login@email.com",
			Password: "login1234",
		}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, response.Code)

		})
	})
}

func TestUpdateUser(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		repository.InsertUser(&entities.User{ID: 1, Login: "user123", Email: "login123@email.com",
			Password: "123456"})

		rest.PUT("/users/1").SetJSONInterface(requests.UserRequest{
			Login:    "userUpdated123",
			Email:    "loginUpdated@email.com",
			Password: "login1234",
		}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, response.Code)

			user, _ := repository.FindUserById(1)

			assert.Equal(t, "userUpdated123", user.Login)
			assert.Equal(t, "loginUpdated@email.com", user.Email)
			assert.Equal(t, "login1234", user.Password)
			assert.NotEmpty(t, user.CreatedAt)
			assert.NotEmpty(t, user.UpdateAt)
		})
	})
}

func TestDeleteUser(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		repository.InsertUser(&entities.User{ID: 1, Login: "user123", Email: "login123@email.com",
			Password: "123456"})

		rest.DELETE("/users/1").
			Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, response.Code)

				planet, _ := repository.FindUserById(1)

				assert.Empty(t, planet)
			})
	})
}
