package controllers

import (
	"net/http"

	dto "user/dtos/user"
	service "user/services/user"

	"github.com/gin-gonic/gin"
)

func InsertUser(c *gin.Context) {
	var user dto.UserInsertDto

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	userToInsert, err := service.UserService.InsertUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, userToInsert)
}

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

func UpdateUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	var user dto.UserUpdateDto

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	userToUpdate, err := service.UserService.UpdateUser(auth, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, userToUpdate)
}

func DeleteUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	userDto, err := service.UserService.DeleteUser(auth)

	if err != nil {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	c.JSON(http.StatusOK, userDto)
}
