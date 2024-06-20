package main

import (
	"log"

	"github.com/streadway/amqp"
)

func ConsumeMessages(ch *amqp.Channel, q amqp.Queue, consumerID int) {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	for d := range msgs {
		log.Printf("Consumer %d received a message: %s", consumerID, d.Body)
	}
}
