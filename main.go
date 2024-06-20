package main

import (
	"log"
)

func main() {

	conn, ch, q := ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	for i := 0; i < 2; i++ {
		publishMessage(ch, q, "Hello World!")
	}

	numConsumers := 3
	for i := 1; i <= numConsumers; i++ {
		go ConsumeMessages(ch, q, i)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}

}
