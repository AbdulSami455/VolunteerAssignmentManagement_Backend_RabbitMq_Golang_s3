package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {

	fmt.Println("GO RabbitMq Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

}
