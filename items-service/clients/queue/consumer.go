package queue

import (
	"log"

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
		"users-delete-queue", // name
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
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

	go consumeMessages("users-delete-queue", msgs)

	log.Printf(" [*] Waiting for messages. To exit, press CTRL+C")
	select {}
}

func consumeMessages(queue string, msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message from queue '%s': %s", queue, d.Body)

		handleDeleteMessage(d.Body)
	}
}

func handleDeleteMessage(body []byte) {
	log.Println("handle user:", string(body))
	/*id := string(body)

	res, err := getItems(id)



	if err != nil {
		log.Println("Cannot delete item:", err)
		log.Println(res.Body)
	} else {
		log.Println("Item deleted:", res.Body)
	}
	fmt.Println("handle delete message")*/
}

/*func getItems(id string) (*http.Response, error) {
	url := "http://items-service:8090/items/user/" + id
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

func deleteItem(id string) (*http.Response, error) {

	// Perform the delete logic here



	return r, nil
}
*/
