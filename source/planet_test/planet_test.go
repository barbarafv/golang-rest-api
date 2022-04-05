package test

import (
	"app/source/domain/entities"
	_ "app/source/planet_test/fixture"
	"app/source/planet_test/testcontainers"
	"app/source/repository"
	"app/source/routes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func init() {
	router = routes.InitRouter()

}

func TestInsertPlanet(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/planets", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	t.Log(w.Body.String())
}

func TestUpdatePlanet(t *testing.T) {

	w := httptest.NewRecorder()

	planet := entities.Planet{Name: "marte", Climate: "tempered", Land: "florests and mountains", Atmosphere: "Type III"}

	jsonPlanet, _ := json.Marshal(planet)

	reader := strings.NewReader(string(jsonPlanet))

	req, _ := http.NewRequest("PUT", "/planets/1", reader)
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

func TestA(t *testing.T) {
	ctx := context.Background()
	containerResult := testcontainers.SetupMysqlContainer(t, ctx, true)
	defer containerResult.Container.Terminate(ctx)

	t.Logf("%s", containerResult.ConnectionURI())
}
