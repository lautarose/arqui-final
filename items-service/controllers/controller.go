package controllers

import (
	"items/dtos"
	service "items/services"
	e "items/utils/errors"
	"net/http"

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
	var items dtos.ItemsDto
	if err := c.BindJSON(&items); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	items, apiErr := ctrl.service.InsertItems(c.Request.Context(), items)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, items)
}

/*
import (
	"items-api/dtos"
	service "items-api/services"
	"items-api/utils/cache"
	"encoding/json"
	"fmt"
	"net/http"


	"github.com/gin-gonic/gin"
)

func GetItemById(c *gin.Context) {

	id := c.Param("id")

	res := cache.Get(id)

	if res != "" {
		var itemDtoCache dtos.ItemDto
		json.Unmarshal([]byte(res), &itemDtoCache)
		fmt.Println("from cache: " + id)
		c.JSON(http.StatusOK, itemDtoCache)
		return
	}

	fmt.Println("not cache: " + id)
	itemDto, er := service.ItemService.GetItemById(id)
	itemDtoStr, _ := json.Marshal(itemDto)
	cache.Set(itemDto.Id, itemDtoStr)
	fmt.Println("save cache: " + itemDto.Id)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, itemDto)
}

func InsertItems(c *gin.Context) {
	var itemsDto dtos.ItemsDto
	err := c.BindJSON(&itemsDto)

	// Error Parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemsDto, er := service.ItemService.InsertItems(itemsDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	for _, itemDto := range itemsDto {
		itemDtoStr, _ := json.Marshal(itemDto)
		cache.Set(itemDto.Id, itemDtoStr)
		fmt.Println("save cache: " + itemDto.Id)
	}

	c.JSON(http.StatusCreated, itemsDto)
}
*/
