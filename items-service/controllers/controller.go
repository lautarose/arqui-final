package controllers

import (
	"items/dtos"
	service "items/services"
	e "items/utils/errors/errors"
	jwt "items/utils/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetItemById(c *gin.Context) {
	item, apiErr := ctrl.service.GetItemById(c.Request.Context(), c.Param("id"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *Controller) InsertItems(c *gin.Context) {

	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	claims, err := jwt.VerifyToken(auth)

	if err != nil {
		apiErr := e.NewForbiddenApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	id, err := strconv.Atoi(claims.Id)

	if err != nil {
		apiErr := e.NewInternalServerApiError("cannot convert claim", err)
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	var items dtos.ItemsDto

	if err := c.BindJSON(&items); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	var itemsToInsert dtos.ItemsDto
	for _, item := range items {
		item.UserID = id
		itemsToInsert = append(itemsToInsert, item)
	}

	items, apiErr := ctrl.service.InsertItems(c.Request.Context(), itemsToInsert)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, items)
}

func (ctrl *Controller) UpdateItem(c *gin.Context) {

	// Obtener los datos del ítem de la solicitud JSON
	var itemDto dtos.ItemDto
	if err := c.ShouldBindJSON(&itemDto); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	// Verificar el userID del token de autorización
	auth := c.GetHeader("Authorization")
	claims, err := jwt.VerifyToken(auth)
	if err != nil {
		apiErr := e.NewForbiddenApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	userID, err := strconv.Atoi(claims.Id)
	if err != nil {
		apiErr := e.NewInternalServerApiError("cannot convert claim", err)
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	// Obtener el ítem original de la base de datos
	originalItem, apiErr := ctrl.service.GetItemById(c.Request.Context(), itemDto.Id)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	// Verificar si el userID del token coincide con el userID del ítem
	if userID != originalItem.UserID {
		apiErr := e.NewForbiddenApiError("No tienes permiso para modificar este ítem")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	// Llamar al servicio para actualizar el ítem
	updatedItem, apiErr := ctrl.service.UpdateItem(c.Request.Context(), itemDto)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	// Devolver el ítem actualizado en formato JSON
	c.JSON(http.StatusOK, updatedItem)
}
