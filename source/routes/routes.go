package routes

import (
	"app/source/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	bindPlanetRoutes(router)
	middleware.BindMiddlewares(router)

	return router
}
