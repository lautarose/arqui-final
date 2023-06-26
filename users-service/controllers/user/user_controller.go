package controllers

import (
	"net/http"

	service "user/services/user"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	userDto, err := service.UserService.GetUser(auth)

	if err != nil {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	c.JSON(http.StatusOK, userDto)
}
