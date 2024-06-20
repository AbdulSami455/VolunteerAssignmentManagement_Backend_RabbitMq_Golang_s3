package main

import (
	"log"
	"time"
)

func main() {

	conn, ch, q := ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	// Publish a message
	for i := 0; i < 10; i++ {
		publishMessage(ch, q, "Hello World!")

	}
	/*
		numConsumers := 10
		for i := 1; i <= numConsumers; i++ {
			go ConsumeMessages(ch, q, i)
		}
	*/
	// Periodically check the status of the queue
	go func() {
		for {
			CheckQueueStatus(ch, q.Name)
			time.Sleep(10 * time.Second) // Check every 2 seconds
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}

}
