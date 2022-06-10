package main

import (
	"blackgo/game/consumer/handlers"
	"blackgo/game/repository"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var r repository.IGameRepository

func main() {

	BROKER_URL := os.Getenv("BROKER_URL")
	conn, err := amqp.Dial(BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	r = repository.NewGormGameRepository()

	err = handlers.Start(ch, &r)
	failOnError(err, "Failed to start handlers")

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	var forever chan struct{}
	<-forever
}
