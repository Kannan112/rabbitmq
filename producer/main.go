package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("failed connection with rabbitMq server")
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to create channle")
	}
	//closding
	defer conn.Close()
	defer ch.Close()

	q, err := ch.QueueDeclare("standalone", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}
	body := "hello-world"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		log.Fatalf("Failed to publish message,%v", err)
	}
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Println("Message published successfully!")
}
