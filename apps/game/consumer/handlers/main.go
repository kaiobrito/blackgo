package handlers

import (
	"blackgo/engine"
	"blackgo/game/queues"
	"blackgo/game/repository"
	"blackgo/messaging"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var repo repository.IGameRepository
var channel *amqp.Channel

func handleCreateGame(msg amqp.Delivery) {
	fmt.Println("handleCreateGame")

	var game engine.Blackgo
	err := json.Unmarshal(msg.Body, &game)

	if err == nil {
		log.Printf("Received a message: %s", string(msg.Body))
		repo.SaveGame(&game)
	} else {
		log.Printf("Not able to parse message: %s", string(msg.Body))
	}
}

func handleGetById(msg amqp.Delivery) {
	fmt.Println("handleGetById " + msg.CorrelationId)

	var body map[string]string
	json.Unmarshal(msg.Body, &body)

	game := repo.GetGameById(body["ID"])
	response, _ := json.Marshal(game)

	fmt.Println("handleGetById", response)

	if err := messaging.Publish(channel, messaging.Message{
		Exchange:      "",
		RoutingKey:    msg.ReplyTo,
		CorrelationId: msg.CorrelationId,
		Body:          response,
	}); err != nil {
		log.Fatalln(err)
	}
}

func Start(ch *amqp.Channel, r *repository.IGameRepository) error {
	repo = *r
	channel = ch
	manager := messaging.NewManager(
		[]messaging.ConsumerHandler{
			messaging.CreateConsumer(queues.GAMES_CREATE_QUEUE, handleCreateGame),
			messaging.CreateConsumer(queues.GAMES_QUERY_QUEUE, handleGetById),
		},
	)

	return manager.Start(ch)
}
