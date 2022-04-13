package routes

import (
	controllers "app/source/controllers"

	"github.com/gin-gonic/gin"
)

func bindPlanetRoutes(router *gin.Engine) {
	planets := router.Group("/planets")
	planets.GET("", controllers.FindPlanets)
	planets.GET("/:id", controllers.FindPlanetById)
	planets.POST("", controllers.InsertPlanet)
	planets.PUT("/:id", controllers.UpdatePlanet)
	planets.DELETE("/:id", controllers.DeletePlanet)

	users := router.Group("/users")
	users.POST("", controllers.InsertUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.GET("/:id", controllers.FindUserById)
	users.DELETE("/:id", controllers.DeleteUser)

}
