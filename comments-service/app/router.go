package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router  *gin.Engine
	webPort = ":8100"
)

func init() {
	router = gin.Default()

	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowAllOrigins = true // Permitir todos los orígenes
	router.Use(cors.New(config))
}

func StartRoute() {
	MapUrls()
	router.Run(webPort)
}
