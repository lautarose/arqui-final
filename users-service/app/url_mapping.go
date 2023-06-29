package app

import (
	loginController "user/controllers/login"
	userController "user/controllers/user"
)

// MapUrls maps the urls
func MapUrls() {

	// Users Mapping
	router.GET("/user", userController.GetUser)
	router.POST("/user/signup", userController.InsertUser)
	router.PUT("/user", userController.UpdateUser)
	router.DELETE("/user", userController.DeleteUser)

	// Login Mapping
	router.POST("/user/login", loginController.Login)
}
