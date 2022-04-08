package test

import (
	_ "app/source/_testinit/fixture"
	"app/source/configuration"
	"app/source/controllers/requests"
	"app/source/domain/entities"
	"app/source/planet_test/testcontainers"
	"app/source/planet_test/testutils"
	"app/source/repository"
	"app/source/routes"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/appleboy/gofight/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var router *gin.Engine
var mysqlContainer testcontainers.ContainerResult
var db *gorm.DB

var RunTest = testutils.CreateForEach(setUp, tearDown)

func TestMain(m *testing.M) {
	log.Println("Starting test setup")
	start := time.Now()
	BeforeAll()
	log.Printf("Setup took %s seconds\n", time.Since(start))
	exitVal := m.Run()
	os.Exit(exitVal)
}

func BeforeAll() {
	router = routes.InitRouter()
	mysqlContainer = testcontainers.SetupMysqlContainer(
		&testcontainers.Testcontainer{
			Database:     configuration.Config.DBName,
			RootPassword: configuration.Config.DBPass,
		},
	)
	repository.OpenConnectionDb()
}

func setUp() {
	repository.AutoMigrate()
}

func tearDown() {
	repository.DropAll()
}
func TestInsertPlanet(t *testing.T) {
	RunTest(func() {
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

	RunTest(func() {
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

	RunTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.POST("/planets").
			SetJSONInterface(requests.PlanetRequest{
				Name:       "",
				Climate:    "random",
				Land:       "random",
				Atmosphere: "random",
			}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusInternalServerError, response.Code)
		})
	})
}

func TestUpdatePlanet(t *testing.T) {

	RunTest(func() {
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

	RunTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.PUT("/planets/1").
			SetJSONInterface(requests.PlanetRequest{
				Name:       "marte",
				Climate:    "tempered",
				Land:       "florests and mountains",
				Atmosphere: "Type III",
			}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusInternalServerError, response.Code)
		})
	})

}
func TestGetPlanet(t *testing.T) {

	RunTest(func() {
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

	RunTest(func() {
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

	RunTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.GET("/planets/1").
			Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, response.Code)
			})
	})
}

func TestDeletePlanet(t *testing.T) {

	RunTest(func() {
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

func TestDeleteNotExistPlanet(t *testing.T) {

	RunTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.DELETE("/planets/1").
			Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, response.Code)
			})
	})

}
