package main

import (
	"comments/app"
	"comments/database"
)

func main() {
	database.StartDbEngine()
	app.StartRoute()
}
