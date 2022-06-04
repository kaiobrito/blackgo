package messaging

import (
	"errors"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func PublishAndConsume(ch *amqp091.Channel, message Message) ([]byte, error) {
	replyQueue := message.ReplyTo
	if replyQueue == nil {
		return nil, errors.New("ReplyTo not defined")
	}

	c := make(chan []byte)

	go func(q Queue) {
		msg, err := SubscribeToQueue(ch, q, amqp091.Table{
			"CorrelationId": message.CorrelationId,
		})

		if err != nil {
			fmt.Println("Error on subscribing:"+q.FullPath(), err)
			close(c)
		}
		data := <-msg
		c <- data.Body
	}(*replyQueue)

	if err := Publish(ch, message); err != nil {
		return nil, err
	}

	return <-c, nil

}

func Publish(ch *amqp091.Channel, message Message) error {
	var replyTo string
	if message.ReplyTo != nil {
		replyTo = message.ReplyTo.Name
	}

	return ch.Publish(
		message.Queue.Exchange.Name, // exchange
		message.Queue.FullPath(),    // routing key
		false,                       // mandatory
		false,                       // immediate
		amqp091.Publishing{
			ContentType:   "text/plain",
			ReplyTo:       replyTo,
			CorrelationId: message.CorrelationId,
			Body:          message.Body,
		})
}
