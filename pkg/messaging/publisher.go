package messaging

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func GetExclusiveQueue(ch *amqp091.Channel) (amqp091.Queue, error) {
	queue := Queue{
		Name:       "",
		RoutingKey: "",
		Exchange: Exchange{
			Name:    "",
			EType:   "",
			Durable: false,
		},
		Durable:   false,
		Exclusive: true,
	}
	return configureQueue(ch, queue)
}

func PublishAndConsume(ch *amqp091.Channel, message Message) ([]byte, error) {
	c := make(chan []byte)

	replyQueue, err := GetExclusiveQueue(ch)

	if err != nil {
		fmt.Println("Error to create Queue", err)
		close(c)
	}
	defer ch.QueueDelete(replyQueue.Name, false, false, true)

	msgs, queue, err := CreateSubscription(ch, replyQueue, amqp091.Table{
		"CorrelationId": message.CorrelationId,
	})
	message.ReplyToName = &replyQueue.Name

	if err != nil {
		fmt.Println("Error on subscribing:"+queue.Name, err)
		close(c)
	}

	if err := Publish(ch, message); err != nil {
		return nil, err
	}

	go func() {
		fmt.Println("Waiting for response")
		for msg := range msgs {
			fmt.Println("Message received", msg.Body)
			if msg.CorrelationId == message.CorrelationId {
				c <- msg.Body
				break
			} else {
				fmt.Println("Different correlationId. Requeuing", msg.CorrelationId)
			}
		}
		fmt.Println("Finish publish and consume")
	}()

	return <-c, nil

}

func Publish(ch *amqp091.Channel, message Message) error {
	var replyTo string
	if message.ReplyToName != nil {
		replyTo = *message.ReplyToName
	}

	log.Default().Printf("Publishing to %s with Correlation %s", message.RoutingKey, message.CorrelationId)

	return ch.Publish(
		message.Exchange,   // exchange
		message.RoutingKey, // routing key
		false,              // mandatory
		false,              // immediate
		amqp091.Publishing{
			ContentType:   "text/plain",
			ReplyTo:       replyTo,
			CorrelationId: message.CorrelationId,
			Body:          message.Body,
		})
}
