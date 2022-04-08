package middleware

import (
	"app/source/domain/exception"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindMiddlewares(router *gin.Engine) {
	log.Println("BindMiddlewares")
	router.Use(gin.CustomRecovery(exceptionMiddleware))
}

func exceptionMiddleware(c *gin.Context, recovered interface{}) {
	except, ok := recovered.(*exception.HttpException)

	if ok {
		c.String(except.StatusCode, except.Message)
	} else {
		log.Printf("Exception not mapped: %s", recovered)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
