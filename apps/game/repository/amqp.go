package repository

import (
	"blackgo/engine"
	"blackgo/game/queues"
	"blackgo/messaging"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
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
		CorrelationId: uuid.NewString(),
		Queue:         queues.GAMES_CREATE_QUEUE,
		ReplyTo:       nil,
		Body:          body,
	})
	fmt.Println(err)
	if err != nil {
		return nil
	}
	return &game
}

func (r AMQPGameRepository) SaveGame(game *engine.Blackgo) {

}
func (r AMQPGameRepository) GetGameById(id string) *engine.Blackgo {
	body, err := json.Marshal(map[string]string{
		"ID": id,
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	data, err := messaging.PublishAndConsume(r.ch, messaging.Message{
		CorrelationId: uuid.NewString(),
		Queue:         queues.GAMES_QUERY_QUEUE,
		ReplyTo:       &queues.GAMES_GET_QUEUE,
		Body:          body,
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}
	var game engine.Blackgo
	if err = json.Unmarshal(data, &game); err != nil {
		fmt.Println(err)
		return nil
	}

	return &game
}
func (r AMQPGameRepository) GetAllGames() map[string]*engine.Blackgo {
	return nil
}
func (r AMQPGameRepository) DeleteAll() {

}
