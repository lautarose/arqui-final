package main

import (
	"search/app"
	queue "search/clients/queue"

	_ "github.com/ugorji/go/codec"
)

func main() {
	go queue.Consume()
	app.StartApp()
}
