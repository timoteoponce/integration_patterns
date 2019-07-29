package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/streadway/amqp"
)

func main() {
	readCredentials()
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@159.65.220.217:5672", user, pwd))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed on channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("rpc_queue", false, false, false, false, nil)
	failOnError(err, "Failed to connect to queue")
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	failOnError(err, "Failed to connect to consume")
	go func() {
		for d := range msgs {
			log.Println("AAAAAA")
			input := string(d.Body)
			response := sayHello(input)

			ch.Publish("", d.ReplyTo, false, false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(response),
				})
			d.Ack(false)
		}
	}()

	forever := make(chan bool)
	log.Printf(" [*] Waiting RPC requests")
	<-forever
}

func sayHello(name string) string {
	adjectives := []string{"lazy", "glorious", "brilliant", "complex"}
	adj := adjectives[rand.Intn(len(adjectives)-1)]
	return fmt.Sprintf("Hello %s %s!!", adj, name)
}

var user string
var pwd string

func readCredentials() {
	fmt.Print("RabbitMQ username: ")
	fmt.Scanln(&user)
	fmt.Print("RabbitMQ password: ")
	fmt.Scanln(&pwd)

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
