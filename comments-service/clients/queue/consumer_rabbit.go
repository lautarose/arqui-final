package queue

import (
	commentsClient "comments/clients/comments"
	dtos "comments/dtos/user"
	jwtUtils "comments/utils/jwt"
	"encoding/json"
	"log"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Channel *amqp.Channel
}

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
		"users-comments-delete-queue", // name
		false,                         // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
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

	go consumeMessages("users-comments-delete-queue", msgs)

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
	log.Printf("handle delete message")
	//parse body
	var requestBody dtos.UserMessageDto
	err := json.Unmarshal(body, &requestBody)
	if err != nil {
		log.Println("Error parsing request body:", err)
		return
	}

	id := requestBody.Id
	token := requestBody.Token

	claims, err := jwtUtils.VerifyToken(token)

	if err != nil {
		log.Println("error verifying token: ", err)
		return
	}

	newId := strconv.Itoa(id)

	if newId != claims.Id {
		log.Println("not matching ids: ", err)
		return
	}

	err = commentsClient.DeleteCommentsByUserId(id)

	if err != nil {
		log.Println("not comments found: ", err)
		return
	}

	log.Println("comments deleted where user id: ", newId)
}
