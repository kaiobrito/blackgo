package messaging

import (
	"github.com/rabbitmq/amqp091-go"
)

type MessageHandler func(amqp091.Delivery)

type ConsumerHandler struct {
	queue      string
	exchange   string
	routingKey string

	handler MessageHandler
}

func CreateConsumer(queue string, routingKey string, exchange string, handler MessageHandler) ConsumerHandler {
	return ConsumerHandler{
		queue:      queue,
		routingKey: routingKey,
		exchange:   exchange,
		handler:    handler,
	}
}

func (handler ConsumerHandler) bind(ch *amqp091.Channel) error {
	return ch.QueueBind(
		handler.queue,
		handler.routingKey,
		handler.exchange,
		false,
		nil,
	)
}

func (handler ConsumerHandler) start(ch *amqp091.Channel) error {
	if err := handler.bind(ch); err != nil {
		return nil
	}

	msgs, err := ch.Consume(
		handler.queue,      // queue
		handler.routingKey, // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)

	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler.handler(d)
		}
	}()
	return nil
}
