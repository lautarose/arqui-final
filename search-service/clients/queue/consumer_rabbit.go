package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	dtos "search/dtos"

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

	queues := []string{"item-insert-queue", "item-update-queue", "item-delete-queue"}

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
		case "item-update-queue":
			handleUpdateMessage(d.Body)
		case "item-delete-queue":
			handleDeleteMessage(d.Body)
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

func handleUpdateMessage(body []byte) {
	r, err := getItem(string(body))
	if err != nil {
		log.Println("Cannot get item:", err)
		return
	}

	res, err := updateItem(r)
	if err != nil {
		log.Println("Cannot update item:", err)
		log.Println(res.Body)
	} else {
		log.Println("Item updated:", res.Body)
	}

	fmt.Println("handle update message")
}

func handleDeleteMessage(body []byte) {
	/*r, err := getItem(string(body))
	if err != nil {
		log.Println("Cannot get item:", err)
		return
	}

	res, err := deleteItem(r)
	if err != nil {
		log.Println("Cannot delete item:", err)
		log.Println(res.Body)
	} else {
		log.Println("Item deleted:", res.Body)
	}*/
	fmt.Println("handle delete message")
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

func updateItem(r *http.Response) (*http.Response, error) {
	//leer el body de la response de getItem
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error al leer el cuerpo de la respuesta:", err)
		return nil, err
	}

	// Decodificar el cuerpo de la respuesta en una estructura Item
	var item dtos.ItemReponseDto //es un itemResponse por como response la api items
	err = json.Unmarshal(body, &item)
	if err != nil {
		fmt.Println("Error al decodificar el cuerpo de la respuesta:", err)
		return nil, err
	}

	//transformo en la estructura necesaria para solr
	itemUpdateDto := dtos.ItemUpdateDto{
		Id: item.Id,
		Title: dtos.FieldValue{
			Set: item.Title,
		},
		Seller: dtos.FieldValue{
			Set: item.Seller,
		},
		Price: dtos.FieldValue{
			Set: item.Price,
		},
		Currency: dtos.FieldValue{
			Set: item.Currency,
		},
		Picture: dtos.FieldValue{
			Set: item.Picture,
		},
		Description: dtos.FieldValue{
			Set: item.Description,
		},
		State: dtos.FieldValue{
			Set: item.State,
		},
		City: dtos.FieldValue{
			Set: item.City,
		},
		Street: dtos.FieldValue{
			Set: item.Street,
		},
		Number: dtos.FieldValue{
			Set: item.Number,
		},
	}

	//solr me pide que sea un slice.
	var itemsToUpdate dtos.ItemsUpdateDto

	itemsToUpdate = append(itemsToUpdate, itemUpdateDto)

	// Codificar el DTO ItemUpdateDto en JSON
	dtoJSON, err := json.Marshal(itemsToUpdate)
	if err != nil {
		fmt.Println("Error al codificar el DTO en JSON:", err)
		return nil, err
	}

	// transformarlo en un objeto reader
	dtoReader := bytes.NewReader(dtoJSON)

	r, err = http.Post("http://solr:8983/solr/items/update?commit=true", "application/json", dtoReader)

	if err != nil {
		log.Println(err)
		return r, err
	}

	return r, nil
}

/*func deleteItem(r *http.Response) (*http.Response, error) {
	body := r.Body

	// Perform the delete logic here

	if err != nil {
		log.Println(err)
		return r, err
	}

	return r, nil
}*/
