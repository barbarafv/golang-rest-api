package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	bindPlanetRoutes(router)

	return router
}
