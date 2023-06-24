package main

import (
	"user/app"
	"user/database"
)

func main() {
	database.StartDbEngine()
	app.StartRoute()
}