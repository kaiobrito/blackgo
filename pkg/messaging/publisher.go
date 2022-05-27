package messaging

import (
	"github.com/rabbitmq/amqp091-go"
)

// func PublishAndWait(ch *amqp091.Channel, message PublishWaitMessage, c chan []byte) error {
// 	pm := message.Publish
// 	pm.CorrelationId = message.CorrelationId
// 	pm.ReplyTo = message.Receive.RoutingKey

// 	Publish(ch, message.Publish)

// 	err := ch.QueueBind(
// 		message.Receive.ReplyTo,
// 		message.CorrelationId,
// 		message.Receive.Exchange,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	msgs, err := ch.Consume(
// 		message.Receive.Exchange,
// 		message.Receive.ReplyTo,
// 		true,
// 		false, false, false,
// 		nil,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	for msg := range msgs {
// 		if msg.CorrelationId == message.CorrelationId {
// 			fmt.Println(msg.Body)
// 			break
// 		}
// 	}

// 	return nil
// }

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
