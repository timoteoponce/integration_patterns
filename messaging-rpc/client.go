package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/streadway/amqp"
)

func main() {
	readCredentials()
	log.Printf("Result from RPC call %s \n", sayHelloRpc("Timoteo"))
}

var user string
var pwd string

func readCredentials() {
	fmt.Print("RabbitMQ username: ")
	fmt.Scanln(&user)
	fmt.Print("RabbitMQ password: ")
	fmt.Scanln(&pwd)
}

func sayHelloRpc(input string) string {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@159.65.220.217:5672", user, pwd))
	failOnError(err, "Failed to connect to consume")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to connect to consume")
	defer ch.Close()

	replyQueue, err := ch.QueueDeclare("", false, false, false, false, nil)
	failOnError(err, "Failed to connect to consume")
	msgs, err := ch.Consume(replyQueue.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to connect to consume")

	correlationId := randomString(32)
	ch.Publish("", "rpc_queue", false, false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: correlationId,
			ReplyTo:       replyQueue.Name,
			Body:          []byte(input),
		})
	var result string
	for d := range msgs {
		if correlationId == d.CorrelationId {
			result = string(d.Body)
			break
		}
	}
	return result
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
