package handlers

import (
	"blackgo/engine"
	"blackgo/game/repository"
	"blackgo/messaging"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var repo repository.IGameRepository
func handleCreateGame(msg amqp.Delivery) {
	var game engine.Blackgo
	err := json.Unmarshal(msg.Body, &game)

	if err == nil {
		log.Printf("Received a message: %s", string(msg.Body))
	} else {
		log.Printf("Not able to parse message: %s", string(msg.Body))
	}
}


func Start(ch *amqp.Channel, r *repository.IGameRepository) error {
	repo = *r
	manager := messaging.NewManager(
		[]messaging.ConsumerHandler{
			messaging.CreateConsumer(queues.GAMES_CREATE_QUEUE, handleCreateGame),
		},
	)

	return manager.Start(ch)
}
