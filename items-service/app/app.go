package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()

	// Configurar el middleware CORS
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	router.Use(cors.New(config))

	deps := BuildDependencies()
	MapUrls(router, deps)
	_ = router.Run(":8090")
}
