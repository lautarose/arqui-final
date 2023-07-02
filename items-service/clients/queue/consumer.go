package queue

import (
	"encoding/json"
	"fmt"
	"io"
	"items/dtos"
	"log"
	"net/http"
	"strconv"

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

	//parse body
	var requestBody dtos.UserMessageDto
	err := json.Unmarshal(body, &requestBody)
	if err != nil {
		log.Println("Error parsing request body:", err)
		return
	}

	id := requestBody.Id
	token := requestBody.Token

	newId := strconv.Itoa(id)

	res, err := getItems(newId)

	if err != nil {
		log.Println("Cannot get items:", err)
		log.Println(res.Body)
	} else {
		defer res.Body.Close()

		// Read the response body
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			return
		}

		// Convert responseBody to a slice of strings
		var items []string
		err = json.Unmarshal(responseBody, &items)
		if err != nil {
			log.Println("Error parsing response:", err)
			return
		}

		for _, item := range items {
			_, err := deleteItem(item, token)
			if err != nil {
				log.Println("Error deleting item: ", item, " error: ", err)
			}
		}
	}

}

func getItems(id string) (*http.Response, error) {
	url := "http://items-service:8090/items/user/" + id
	r, err := http.Get(url)

	if err != nil {
		log.Panic(err)
		return r, err
	}

	return r, nil
}

func deleteItem(item string, token string) (*http.Response, error) {

	// URL de la solicitud DELETE
	url := fmt.Sprintf("http://localhost:8090/items/%s", item)

	// Crear la solicitud DELETE
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error al crear la solicitud DELETE:", err)
		return nil, err
	}

	// Agregar el token JWT al encabezado de autorización
	req.Header.Set("Authorization", token)

	// Realizar la solicitud HTTP DELETE
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud DELETE:", err)
		return resp, err
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Elemento eliminado exitosamente")
	} else {
		fmt.Println("Error al eliminar el elemento. Código de estado:", resp.StatusCode)
	}
	return resp, nil
}
