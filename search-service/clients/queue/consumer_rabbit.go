package clients

import (
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume() error {
	conn, err := amqp.Dial("amqp://user:password@rabbit:5672/")
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	queues := []string{"item-insert-queue", "item-modification-queue", "item-deletion-queue"}

	for _, queue := range queues {
		q, err := ch.QueueDeclare(
			queue, // name
			false, // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			return err
		}

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		if err != nil {
			return err
		}

		go consumeMessages(queue, msgs)
	}

	log.Printf(" [*] Waiting for messages. To exit, press CTRL+C")
	select {}
}

func consumeMessages(queue string, msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message from queue '%s': %s", queue, d.Body)

		switch queue {
		case "item-insert-queue":
			handleInsertMessage(d.Body)
		case "item-modification-queue":
			handleModificationMessage(d.Body)
		case "item-deletion-queue":
			handleDeletionMessage(d.Body)
		}
	}
}

func handleInsertMessage(body []byte) {
	r, err := getItem(string(body))
	if err != nil {
		log.Println("Cannot get item:", err)
		return
	}

	res, err := insertItem(r)
	if err != nil {
		log.Println("Cannot insert item:", err)
		log.Println(res.Body)
	} else {
		log.Println("Item inserted:", res.Body)
	}
}

func handleModificationMessage(body []byte) {
	// Handle modification message logic here
	log.Println("Received a modification message:", string(body))
}

func handleDeletionMessage(body []byte) {
	// Handle deletion message logic here
	log.Println("Received a deletion message:", string(body))
}

func getItem(id string) (*http.Response, error) {
	url := "http://items-service:8090/items/" + id
	r, err := http.Get(url)

	if err != nil {
		log.Panic(err)
		return r, err
	}

	return r, nil
}

func insertItem(r *http.Response) (*http.Response, error) {
	body := r.Body

	r, err := http.Post("http://solr:8983/solr/items/update/json/docs?commit=true", "application/json", body)

	if err != nil {
		log.Println(err)
		return r, err
	}

	return r, nil
}
