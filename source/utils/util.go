package utils

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadBody(c *gin.Context, request any) {
	err := c.BindJSON(&request)

	if err != nil {
		log.Panic("<readBody> Error to bind JSON", err)
	}
}

func ConvertToInt(stringValue string) int {
	valueConv, err := strconv.Atoi(stringValue)

	if err != nil {
		log.Panic("<ConvertToString> Error to convert to int", err)
	}

	return valueConv
}

func ConvertToString(intValue int) string {
	valueConv := strconv.Itoa(intValue)

	return valueConv
}
