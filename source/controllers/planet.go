package controllers

import (
	"app/source/controllers/requests"
	"app/source/service"
	"app/source/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPlanets(c *gin.Context) {
	result := service.FindPlanets()
	c.JSON(http.StatusOK, result)
}

func FindPlanetById(c *gin.Context) {
	result := service.FindPlanetById(utils.ConvertToInt(c.Params.ByName("id")))
	c.JSON(http.StatusOK, result)
}

func UpdatePlanet(c *gin.Context) {
	updatePlanetRequest := requests.PlanetRequest{}

	id := c.Params.ByName("id")

	utils.ReadBody(c, &updatePlanetRequest)

	service.UpdatePlanet(&updatePlanetRequest, utils.ConvertToInt(id))
	c.JSON(http.StatusOK, gin.H{"id" + id: "is updated"})
}

func DeletePlanet(c *gin.Context) {

	id := c.Params.ByName("id")

	service.DeletePlanet(utils.ConvertToInt(id))
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}

func InsertPlanet(c *gin.Context) {
	planet := requests.PlanetRequest{}

	utils.ReadBody(c, &planet)
	service.InsertPlanet(&planet)
	c.JSON(http.StatusOK, planet)
}
