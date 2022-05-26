package handlers

import (
	"blackgo/engine"
	"blackgo/messaging"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func handleCreateGame(msg amqp.Delivery) {
	var game engine.Blackgo
	err := json.Unmarshal(msg.Body, &game)

	if err == nil {
		log.Printf("Received a message: %s", string(msg.Body))
	} else {
		log.Printf("Not able to parse message: %s", string(msg.Body))
	}
}

func Start(ch *amqp.Channel) error {
	manager := messaging.NewManager(
		[]messaging.ConsumerHandler{
			messaging.CreateConsumer("games.action", "create", "amqp.fanout", handleCreateGame),
		},
	)

	return manager.Start(ch)
}
