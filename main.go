package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {

	fmt.Println("GO RabbitMq Tutorial")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel:", err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Failed to declare a queue:", err)
		panic(err)
	}

	fmt.Printf("Queue declared: %v\n", q)

	body := "Hello World"
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		fmt.Println("Failed to publish a message:", err)
		panic(err)
	}

	fmt.Printf("Successfully published message: %s\n", body)
}
