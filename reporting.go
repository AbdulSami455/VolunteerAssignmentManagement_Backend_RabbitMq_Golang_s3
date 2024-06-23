package main

import (
	"log"

	"github.com/streadway/amqp"
)

func CheckQueueStatus(ch *amqp.Channel, queueName string) {
	q, err := ch.QueueInspect(queueName)
	FailOnError(err, "Failed to inspect queue")

	log.Printf("Queue %s has %d messages and %d consumers", q.Name, q.Messages, q.Consumers)
}
