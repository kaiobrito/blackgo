package repository

import (
	"blackgo/engine"
	"blackgo/messaging"
	"encoding/json"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type AMQPGameRepository struct {
	ch *amqp091.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewAMQPGameRepository() IGameRepository {
	BROKER_URL := os.Getenv("BROKER_URL")
	conn, err := amqp091.Dial(BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")

	return &AMQPGameRepository{
		ch: ch,
	}
}

func (r AMQPGameRepository) CreateGame() *engine.Blackgo {
	game := engine.NewBlackgoGameWithShuffler(engine.DefaultShuffler())
	game.Start()
	body, err := json.Marshal(game)

	if err != nil {
		return nil
	}

	err = messaging.Publish(r.ch, messaging.Message{
		RoutingKey: "game.actions.create",
		Exchange:   "amq.fanout",
		Body:       body,
	})
	if err != nil {
		return nil
	}
	return &game
}

func (r AMQPGameRepository) SaveGame(game *engine.Blackgo) {

}
func (r AMQPGameRepository) GetGameById(id string) *engine.Blackgo {
	return nil
}
func (r AMQPGameRepository) GetAllGames() map[string]*engine.Blackgo {
	return nil
}
func (r AMQPGameRepository) DeleteAll() {

}
