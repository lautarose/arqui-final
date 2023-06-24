package app

import (
	controllers "search/controllers"
	service "search/services"
)

type Dependencies struct {
	ItemController *controllers.Controller
}

func BuildDependencies() *Dependencies {

	// Services
	service := service.NewServiceImpl()

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		ItemController: controller,
	}
}
