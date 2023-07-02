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
	q, err := queue.Channel.QueueDeclare(
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

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body := msg

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	err = queue.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // inmediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(bodyBytes),
		}, //message
	)
	if err != nil {
		return err
	}

	b := strconv.Itoa(body.Id)

	log.Printf(" [RabbitMQ] Sent %s user-delete-queue", b)

	return nil
}
