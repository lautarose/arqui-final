package controllers

import (
	"net/http"
	"strconv"

	service "user/services/user"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, userDto)
		return
	}

	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, usersDto)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}
