package middleware

import (
	"app/source/domain/exception"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindMiddlewares(router *gin.Engine) {
	router.Use(gin.CustomRecovery(exceptionMiddleware))
}

func exceptionMiddleware(c *gin.Context, recovered interface{}) {
	if except, ok := recovered.(exception.HttpException); ok {
		c.String(except.StatusCode, except.Message)
	}
	log.Println("Exception not mapped")
	c.AbortWithStatus(http.StatusInternalServerError)
}
