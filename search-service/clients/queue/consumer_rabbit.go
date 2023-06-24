package clients

import (
	"encoding/json"
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

	q, err := ch.QueueDeclare(
		"task_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
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

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			r, err := getItem(string(d.Body))

			if err != nil {
				log.Println("cannot get item")
			}

			res, err := insertItem(r)

			if err != nil {
				log.Print("cannot insert item")
				log.Println(res.Body)
			} else {
				log.Println("item inserted")
				log.Println(res.Body)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
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

func ParseBody(bytes []byte) (string, error) {
	var id string
	if err := json.Unmarshal(bytes, &id); err != nil {
		log.Println(err)
		return "0", err
	}
	return id, nil
}
