package controllers

import (
	"aplicacao/dto/requests"
	"aplicacao/source/domain/entities"
	"aplicacao/source/service"
	"aplicacao/source/utils"
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
	updatePlanetRequest := requests.UpdatePlanetRequest{}

	id := c.Params.ByName("id")
	utils.ReadBody(c, &updatePlanetRequest)

	service.UpdatePlanet(&updatePlanetRequest, id)
	c.JSON(http.StatusOK, updatePlanetRequest)
}

func DeletePlanet(c *gin.Context) {

	id := c.Params.ByName("id")

	service.DeletePlanet(id)
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}

func InsertPlanet(c *gin.Context) {
	planet := entities.Planet{}

	utils.ReadBody(c, &planet)
	service.InsertPlanet(&planet)
	c.JSON(http.StatusOK, planet)
}
