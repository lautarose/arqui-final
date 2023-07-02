package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine, dependencies *Dependencies) {
	// Products Mapping

	router.GET("/items/:id", dependencies.ItemController.GetItemById)
	router.GET("/items/user/:id", dependencies.ItemController.GetItemsIdByUserId)
	router.POST("/items/load", dependencies.ItemController.InsertItems)
	router.PUT("/items/update", dependencies.ItemController.UpdateItem)
	router.DELETE("/items/:id", dependencies.ItemController.DeleteItem)

	fmt.Println("Finishing mappings configurations")
}
