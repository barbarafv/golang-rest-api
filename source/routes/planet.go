package routes

import (
	controllers "app/source/controllers/planets"

	"github.com/gin-gonic/gin"
)

func bindPlanetRoutes(router *gin.Engine) {
	planets := router.Group("/planets")
	planets.GET("", controllers.FindPlanets)
	planets.GET("/:id", controllers.FindPlanetById)
	planets.POST("", controllers.InsertPlanet)
	planets.PUT("/:id", controllers.UpdatePlanet)
	planets.DELETE("/:id", controllers.DeletePlanet)
}
