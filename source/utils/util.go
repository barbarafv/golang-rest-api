package utils

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadBody(c *gin.Context, entity any) {
	err := c.BindJSON(&entity)

	if err != nil {
		log.Panic("<readBody> Error to bind JSON", err)
	}
}

func ConvertToString(stringValue string) int {
	valueConv, err := strconv.Atoi(stringValue)

	if err != nil {
		log.Panic("<ConvertToString> Error to convert to int", err)
	}

	return valueConv
}
