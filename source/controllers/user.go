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

func UpdateUser(c *gin.Context) {
	userRequest := requests.UserRequest{}

	id := c.Params.ByName("id")
	utils.ReadBody(c, &userRequest)

	service.UpdateUser(&userRequest, utils.ConvertToInt(id))
	c.JSON(http.StatusOK, gin.H{"user " + id: "is updated"})
}

func FindUserById(c *gin.Context) {
	result := service.FindUserById(utils.ConvertToInt(c.Params.ByName("id")))
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	service.DeleteUser(utils.ConvertToInt(id))
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}
