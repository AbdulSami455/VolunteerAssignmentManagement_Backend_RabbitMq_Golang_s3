package main

import (
	"log"
)

func main() {

	conn, ch, q := ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	numConsumers := 3
	for i := 1; i <= numConsumers; i++ {
		go ConsumeMessages(ch, q, i)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}

}
