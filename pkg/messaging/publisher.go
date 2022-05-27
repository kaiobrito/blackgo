package messaging

import (
	"github.com/rabbitmq/amqp091-go"
)

func PublishAndConsume(ch *amqp091.Channel, message Message) ([]byte, error) {
	Publish(ch, message)
	if message.ReplyTo != nil {
		c := make(chan []byte)
		go func() {
			msg, err := SubscribeToQueue(ch, *message.ReplyTo, amqp091.Table{
				"CorrelationId": message.CorrelationId,
			})

			if err != nil {
				close(c)
			}
			data := <-msg
			c <- data.Body
		}()

		return <-c, nil
	}

	return nil, nil
}

func Publish(ch *amqp091.Channel, message Message) error {
	replyTo := ""
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
