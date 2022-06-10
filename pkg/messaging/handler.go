package messaging

import (
	"fmt"

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

func SubscribeToQueue(ch *amqp091.Channel, q Queue, table amqp091.Table) (<-chan amqp091.Delivery, error) {
	correlationId := table["CorrelationId"]
	fmt.Printf("Subscribing to %s with CorrelationID: %v\n", q.FullPath(), correlationId)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		table,  // args
	)

	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (handler ConsumerHandler) start(ch *amqp091.Channel) error {
	go func() {

		msgs, err := SubscribeToQueue(ch, handler.queue, nil)
		if err != nil {
			panic(err)
		}
		for msg := range msgs {
			handler.handler(msg)
		}
		fmt.Println("Stopping handler " + handler.queue.Name)
	}()

	return nil
}
