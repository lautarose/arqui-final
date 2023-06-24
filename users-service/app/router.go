package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	webPort = ":8081" 
)

func init() {
	router = gin.Default()
	router.Use(cors.Default())
}

func StartRoute() {
	MapUrls()
	router.Run(webPort)
}
