package messaging

import (
	"github.com/rabbitmq/amqp091-go"
)

type MessageHandler func(amqp091.Delivery)

type ConsumerHandler struct {
	queue Queue

	handler MessageHandler
}

func CreateConsumer(queue Queue, handler MessageHandler) ConsumerHandler {
	return ConsumerHandler{
		queue:   queue,
		handler: handler,
	}
}

func (handler ConsumerHandler) start(ch *amqp091.Channel) error {
	msgs, err := ch.Consume(
		handler.queue.Name,       // queue
		handler.queue.RoutingKey, // consumer
		true,                     // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)

	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			handler.handler(msg)
		}
	}()

	return nil
}
