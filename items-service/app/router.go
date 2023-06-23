package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine, dependencies *Dependencies) {
	// Products Mapping

	router.GET("/items/:id", dependencies.ItemController.GetItemById)
	router.POST("/items/load", dependencies.ItemController.InsertItem)

	fmt.Println("Finishing mappings configurations")
}
