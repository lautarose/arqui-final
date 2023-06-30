package controllers

import (
	"net/http"
	service "search/services"

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

func (ctrl *Controller) GetItemsByQuery(c *gin.Context) {
	item, apiErr := ctrl.service.GetItemsByQuery(c.Param("query"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *Controller) GetItems(c *gin.Context) {
	item, apiErr := ctrl.service.GetItems()
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}
