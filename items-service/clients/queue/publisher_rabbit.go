package queue

import (
	"items/dtos"
	"strconv"

	"context"
	"fmt"
	"log"
	"time"

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

func (queue RabbitMQ) PublishInsert(ctx context.Context, item dtos.ItemDto) error {
	q, err := queue.Channel.QueueDeclare(
		"item-insert-queue", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body := item.Id
	err = queue.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // inmediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		}, //message
	)
	if err != nil {
		return err
	}

	log.Printf(" [RabbitMQ] Sent %s item-insert-queue", body)

	return nil
}

func (queue RabbitMQ) PublishDelete(ctx context.Context, id string) error {
	q, err := queue.Channel.QueueDeclare(
		"item-delete-queue", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body := id
	err = queue.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // inmediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		}, //message
	)
	if err != nil {
		return err
	}

	log.Printf(" [RabbitMQ] Sent %s from item-delete-queue", body)

	return nil
}

func (queue RabbitMQ) PublishUpdate(ctx context.Context, item dtos.ItemDto) error {
	q, err := queue.Channel.QueueDeclare(
		"item-update-queue", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body := item.Id
	err = queue.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // inmediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		}, //message
	)
	if err != nil {
		return err
	}

	log.Printf(" [RabbitMQ] Sent %s from item-update-queue", body)

	return nil
}
