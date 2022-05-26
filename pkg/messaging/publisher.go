package messaging

import "github.com/rabbitmq/amqp091-go"

type Message struct {
	RoutingKey string
	Exchange   string
	Body       []byte
}

func Publish(ch *amqp091.Channel, message Message) error {
	return ch.Publish(
		message.Exchange,   // exchange
		message.RoutingKey, // routing key
		false,              // mandatory
		false,              // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        message.Body,
		})
}
