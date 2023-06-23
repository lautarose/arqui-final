package app

import (
	"items-service/clients/queue"
	controllers "items-service/controllers"
	service "items-service/services"
	"items-service/services/repositories"
	"time"
)

type Dependencies struct {
	ItemController *controllers.Controller
}

func BuildDependencies() *Dependencies {
	// Repositories
	ccache := repositories.NewCCache(1000, 100, 30*time.Second)
	memcached := repositories.NewMemcached("memcached", 11211)
	mongo := repositories.NewMongoDB("mongo", 27017, "items")
	rabbit := queue.NewRabbitmq("rabbit", 5672)

	// Services
	service := service.NewServiceImpl(ccache, memcached, mongo, rabbit)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		ItemController: controller,
	}
}