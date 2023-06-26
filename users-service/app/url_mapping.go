package app

import (
	loginController "user/controllers/login"
	userController "user/controllers/user"
)

// MapUrls maps the urls
func MapUrls() {

	// Users Mapping
	router.GET("/user", userController.GetUser)

	// Login Mapping
	router.POST("/user/login", loginController.Login)
}
