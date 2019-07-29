package main

import (
	"log"

	"github.com/streadway/amqp"
)

// -connect to rabbitmq broker
// -connect to a queue
// -listen for messages coming to the queue
// -program waiting for more messages to come
func main() {
	conn, err := amqp.Dial("amqp://admin:Password123@159.65.220.217:5672")
	failOnError(err, "Failed to connect")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare("hello-queue", false, false, false, false, nil)
	failOnError(err, "Failed to create a queue")

	msgs, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	failOnError(err, "Failed to consume queue messages ")

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)
			d.Ack(true)
		}
	}()
	log.Printf("[*] Waiting for messages, please exit the program to stop")
	forever := make(chan bool)
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
