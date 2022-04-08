package utils

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadBody(c *gin.Context, request any) {
	err := c.ShouldBindJSON(&request)

	if err != nil {
		log.Panic("<readBody> Error to bind JSON", err)
	}
}

func ConvertToInt(stringValue string) int {
	intValue, err := strconv.Atoi(stringValue)

	if err != nil {
		log.Panic("<ConvertToInt> error to convert parameter to int")
	}

	return intValue
}
