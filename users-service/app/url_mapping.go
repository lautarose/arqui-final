package app

import userController "user/controllers"

// MapUrls maps the urls
func MapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
}