package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine, dependencies *Dependencies) {
	// Products Mapping

	router.GET("/search/", dependencies.ItemController.GetItems)
	router.GET("/search/:query", dependencies.ItemController.GetItemsByQuery)

	fmt.Println("Finishing mappings configurations")
}
