package controllers

import (
	jwtUtils "comments/utils/jwt"
	"net/http"

	dto "comments/dtos/comment"
	service "comments/services/comment"

	"github.com/gin-gonic/gin"
)

func InsertComment(c *gin.Context) {
	var comment dto.CommentInsertDto

	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	claims, err := jwtUtils.VerifyToken(auth)

	if err != nil {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	commentToInsert, err := service.CommentService.InsertComment(claims.Id, comment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, commentToInsert)
}

func GetComments(c *gin.Context) {

	commentsDto, err := service.CommentService.GetComments(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	c.JSON(http.StatusOK, commentsDto)
}

func DeleteComment(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	userDto, err := service.CommentService.DeleteComment(auth, c.Param("id"))

	if err != nil {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	c.JSON(http.StatusOK, userDto)
}
