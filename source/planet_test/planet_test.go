package test

import (
	_ "app/source/_testinit/fixture"
	"app/source/configuration"
	"app/source/planet_test/testcontainers"
	"app/source/repository"
	"app/source/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var mysqlContainer testcontainers.ContainerResult

func TestInsertPlanet(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/planets", strings.NewReader(`{
        "id": "1",
        "name": "marte",
        "climate": "random",
        "land": "random",
        "atmosphere": "random"
    }`))

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	t.Log(w.Body.String())
}

func TestUpdatePlanet(t *testing.T) {

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/planets/1", strings.NewReader(`{
        "name": "marte",
        "climate": "tempered",
        "land": "florests and mountains",
        "atmosphere": "Type III"
    }`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	planetAfterUpdate, _ := repository.FindPlanetById(1)

	assert.Equal(t, "marte", planetAfterUpdate.Name)
	assert.Equal(t, "tempered", planetAfterUpdate.Climate)
	assert.Equal(t, "florests and mountains", planetAfterUpdate.Land)
	assert.Equal(t, "Type III", planetAfterUpdate.Atmosphere)

}
func TestGetPlanet(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/planets", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetPlanetById(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/planets/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	t.Log(w.Body.String())
}

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
	db := repository.OpenConnectionDb()

	query := `INSERT INTO planet (id, name, climate,land, atmosphere)
			  VALUES (1, 'marte', 'random','random','random'),
			  		 (2, venus, random,random,random)`

	db.DB().Exec(query)

}
