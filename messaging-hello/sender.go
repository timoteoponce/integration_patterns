package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// -connect to rabbitmq broker
// -create a queue
// -send a message to the queue
// -end the program
func main() {
	var user string
	var pwd string
	fmt.Print("RabbitMQ username: ")
	fmt.Scanln(&user)
	fmt.Print("RabbitMQ password: ")
	fmt.Scanln(&pwd)

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@159.65.220.217:5672", user, pwd))
	failOnError(err, "Failed to connect")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare("hello-queue", false, false, false, false, nil)
	failOnError(err, "Failed to create a queue")

	for i := 0; i < 10000; i++ {
		msg := fmt.Sprintf("%d Hello from a messaging sender!!!", i)
		ch.Publish("", queue.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	}

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
