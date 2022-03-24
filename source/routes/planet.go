package routes

import (
	controllers "aplicacao/source/controllers/planets"

	"github.com/gin-gonic/gin"
)

func bindPlanetRoutes(router *gin.Engine) {
	router.GET("/planets", controllers.FindPlanets)
	router.GET("/planets/:id", controllers.FindPlanetById)
	router.POST("/planets", controllers.InsertPlanet)
	router.DELETE("/planets/:id", controllers.DeletePlanet)
	router.PUT("/planets/:id", controllers.UpdatePlanet)
}
