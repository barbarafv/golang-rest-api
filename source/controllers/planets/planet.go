package controllers

import (
	"aplicacao/source/domain/entity"
	"aplicacao/source/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPlanets(c *gin.Context) {
	result := service.FindPlanets()
	c.JSON(http.StatusOK, result)
}

func FindPlanetById(c *gin.Context) {
	id := c.Params.ByName("id")

	result := service.FindPlanetById(id)
	c.JSON(http.StatusOK, result)
}

func UpdatePlanet(c *gin.Context) {
	planet := entity.Planet{}

	id := c.Params.ByName("id")
	readBody(c, &planet)

	service.UpdatePlanet(&planet, id)
	c.JSON(http.StatusOK, planet)
}

func DeletePlanet(c *gin.Context) {

	id := c.Params.ByName("id")

	service.DeletePlanet(id)
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}

func InsertPlanet(c *gin.Context) {
	planet := entity.Planet{}

	readBody(c, &planet)
	service.InsertPlanet(&planet)
	c.JSON(http.StatusOK, planet)
}

func readBody(c *gin.Context, entity any) {
	err := c.BindJSON(&entity)

	if err != nil {
		log.Panic("<readBody> Error to bind JSON")
	}
}
