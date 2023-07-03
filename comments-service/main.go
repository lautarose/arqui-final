package main

import (
	"comments/app"
	"comments/clients/queue"
	"comments/database"
)

func main() {
	database.StartDbEngine()
	go queue.Consume()
	app.StartRoute()
}
