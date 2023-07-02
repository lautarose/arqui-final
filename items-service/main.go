package main

import (
	"items/app"
	"items/clients/queue"

	_ "github.com/ugorji/go/codec"
)

func main() {
	go queue.Consume()
	app.StartApp()
}
