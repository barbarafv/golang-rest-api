package utils

import (
	"app/source/domain/exception"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadBody(c *gin.Context, request any) {
	err := c.ShouldBindJSON(&request)

	if err != nil {
		panic(&exception.HttpException{StatusCode: http.StatusBadRequest,
			Message: fmt.Sprint("Error to bind JSON", err)})
	}
}

func ConvertToInt(stringValue string) int {
	intValue, err := strconv.Atoi(stringValue)

	if err != nil {
		panic(&exception.HttpException{StatusCode: http.StatusInternalServerError,
			Message: fmt.Sprint("Error to convert parameter to int", err)})
	}
	return intValue
}
