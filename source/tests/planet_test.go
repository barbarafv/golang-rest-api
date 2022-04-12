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

func TestInsertPlanet(t *testing.T) {
	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.POST("/planets").
			SetJSONInterface(requests.PlanetRequest{
				Name:       "marte",
				Climate:    "random",
				Land:       "random",
				Atmosphere: "random",
			}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, response.Code)

			existPlanet := repository.ExistsPlanetByName("marte")

			assert.True(t, existPlanet)
		})
	})
}

func TestInsertPlanetThatAreadyExist(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		repository.InsertPlanet(&entities.Planet{Id: 1, Name: "marte", Land: "random",
			Climate: "random", Atmosphere: "random"})

		rest.POST("/planets").
			SetJSONInterface(requests.PlanetRequest{
				Name:       "marte",
				Climate:    "random",
				Land:       "random",
				Atmosphere: "random",
			}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, response.Code)
		})
	})
}

func TestInsertPlanetWithoutName(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.POST("/planets").
			SetJSONInterface(requests.PlanetRequest{
				Name:       "",
				Climate:    "random",
				Land:       "random",
				Atmosphere: "random",
			}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, response.Code)
		})
	})
}

func TestUpdatePlanet(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		repository.InsertPlanet(&entities.Planet{Id: 1, Name: "marte", Land: "random",
			Climate: "random", Atmosphere: "random"})

		rest.PUT("/planets/1").
			SetJSONInterface(requests.PlanetRequest{
				Name:       "marte",
				Climate:    "tempered",
				Land:       "florests and mountains",
				Atmosphere: "Type III",
			}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, response.Code)

			planet, _ := repository.FindPlanetById(1)

			assert.Equal(t, "marte", planet.Name)
			assert.Equal(t, "tempered", planet.Climate)
			assert.Equal(t, "florests and mountains", planet.Land)
			assert.Equal(t, "Type III", planet.Atmosphere)
		})
	})
}

func TestUpdatePlanetThatNotExist(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.PUT("/planets/1").
			SetJSONInterface(requests.PlanetRequest{
				Name:       "marte",
				Climate:    "tempered",
				Land:       "florests and mountains",
				Atmosphere: "Type III",
			}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, response.Code)
		})
	})

}
func TestGetPlanet(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		repository.InsertPlanet(&entities.Planet{Id: 1, Name: "marte", Land: "random",
			Climate: "random", Atmosphere: "random"})

		rest.GET("/planets").
			Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, response.Code)

				assert.Equal(t, `[{"id":"1","name":"marte","climate":"random","land":"random","atmosphere":"random"}]`, response.Body.String())
			})
	})
}

func TestGetPlanetById(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		repository.InsertPlanet(&entities.Planet{Id: 1, Name: "marte", Land: "random",
			Climate: "random", Atmosphere: "random"})

		rest.GET("/planets/1").
			Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, response.Code)

				assert.Equal(t, `{"id":"1","name":"marte","climate":"random","land":"random","atmosphere":"random"}`, response.Body.String())
			})
	})
}

func TestGetPlanetNotFound(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.GET("/planets/1").
			Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, response.Code)
			})
	})
}

func TestDeletePlanet(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		repository.InsertPlanet(&entities.Planet{Id: 1, Name: "marte", Land: "random",
			Climate: "random", Atmosphere: "random"})

		rest.DELETE("/planets/1").
			Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, response.Code)

				planet, _ := repository.FindPlanetById(1)

				assert.Empty(t, planet)
			})
	})
}

func TestDeletePlanetNotExist(t *testing.T) {

	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.DELETE("/planets/1").
			Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, response.Code)
			})
	})

}
