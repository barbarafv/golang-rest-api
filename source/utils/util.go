package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ReadBody(c *gin.Context, request any) {
	err := c.BindJSON(&request)

	if err != nil {
		log.Panic("<readBody> Error to bind JSON", err)
	}
}
