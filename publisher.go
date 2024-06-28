package main

import (
	"github.com/streadway/amqp"
)

func publishMessage(ch *amqp.Channel, q amqp.Queue, message string) {
	err := ch.Publish(
		"",
		"hello",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	FailOnError(err, "Failed to publish a message")
}
