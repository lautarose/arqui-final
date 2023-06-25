package app

import userController "user/controllers/user"

// MapUrls maps the urls
func MapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
}
