package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	router.Use(cors.Default())
	deps := BuildDependencies()
	MapUrls(router, deps)
	_ = router.Run(":8090")
}
