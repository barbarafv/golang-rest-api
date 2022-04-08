package routes

import (
	"app/source/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	middleware.BindMiddlewares(router)
	bindPlanetRoutes(router)

	return router
}
