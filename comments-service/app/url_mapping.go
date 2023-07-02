package app

import (
	commentsController "comments/controllers/comments"
)

// MapUrls maps the urls
func MapUrls() {

	// Comments Mapping
	router.GET("/comments/:id", commentsController.GetComments)
	router.POST("/comments/load", commentsController.InsertComment)
	//router.DELETE("/comments/:id")
}
