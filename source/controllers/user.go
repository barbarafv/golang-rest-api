package controllers

import (
	"app/source/controllers/requests"
	"app/source/service"
	"app/source/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertUser(c *gin.Context) {
	user := requests.UserRequest{}

	utils.ReadBody(c, &user)
	service.InsertUser(&user)
	c.JSON(http.StatusOK, user)
}
