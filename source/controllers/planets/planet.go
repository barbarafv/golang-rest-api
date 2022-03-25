package controllers

import (
	"aplicacao/source/dto/requests"
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
	idConv := utils.ConvertToInt(id)

	result := service.FindPlanetById(idConv)
	c.JSON(http.StatusOK, result)
}

func UpdatePlanet(c *gin.Context) {
	updatePlanetRequest := requests.PlanetRequest{}

	id := c.Params.ByName("id")

	utils.ReadBody(c, &updatePlanetRequest)

	idConv := utils.ConvertToInt(id)
	service.UpdatePlanet(&updatePlanetRequest, idConv)
	c.JSON(http.StatusOK, gin.H{"id" + id: "is updated"})
}

func DeletePlanet(c *gin.Context) {

	id := c.Params.ByName("id")
	idConv := utils.ConvertToInt(id)

	service.DeletePlanet(idConv)
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}

func InsertPlanet(c *gin.Context) {
	planet := requests.PlanetRequest{}

	utils.ReadBody(c, &planet)
	service.InsertPlanet(&planet)
	c.JSON(http.StatusOK, planet)
}
