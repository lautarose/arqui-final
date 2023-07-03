package queue

import (
	"encoding/json"
	"strconv"

	"context"
	"fmt"
	"log"
	"time"
	dto "user/dtos/user"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Channel *amqp.Channel
}

func NewRabbitmq(host string, port int) *RabbitMQ {
	portS := strconv.Itoa(port)
	dial := "amqp://user:password@" + host + ":" + portS + "/"
	conn, err := amqp.Dial(dial)
	if err != nil {
		panic(fmt.Sprintf("Error initializing RabbitMQ: %v", err))
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf("Error initializing RabbitMQ: %v", err))
	}

	fmt.Println("[RabbitMQ] Initialized connection")
	return &RabbitMQ{
		Channel: ch,
	}
}

func (queue RabbitMQ) Publish(ctx context.Context, msg dto.UserMessageDto) error {
	// Declarar la primera cola: users-items-delete-queue
	q1, err := queue.Channel.QueueDeclare(
		"users-items-delete-queue", // name
		false,                      // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	if err != nil {
		return err
	}

	// Declarar la segunda cola: users-comments-delete-queue
	q2, err := queue.Channel.QueueDeclare(
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

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body := msg

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Publicar en la primera cola: users-items-delete-queue
	err = queue.Channel.PublishWithContext(ctx,
		"",        // exchange
		q1.Name,   // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(bodyBytes),
		},
	)
	if err != nil {
		return err
	}

	// Publicar en la segunda cola: users-comments-delete-queue
	err = queue.Channel.PublishWithContext(ctx,
		"",        // exchange
		q2.Name,   // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(bodyBytes),
		},
	)
	if err != nil {
		return err
	}

	b := strconv.Itoa(body.Id)

	log.Printf(" [RabbitMQ] Sent %s users-items-delete-queue and users-comments-delete-queue", b)

	return nil
}
