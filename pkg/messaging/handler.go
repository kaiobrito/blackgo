package messaging

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageHandler func(amqp091.Delivery)

type ConsumerHandler struct {
	queue   Queue
	handler MessageHandler
}

func CreateConsumer(queue Queue, handler MessageHandler) ConsumerHandler {
	return ConsumerHandler{
		queue:   queue,
		handler: handler,
	}
}

func declareExchange(ch *amqp.Channel, e Exchange) error {
	log.Default().Println("Creating exchange " + e.Name)
	if err := ch.ExchangeDeclare(e.Name, e.EType, e.Durable, false, false, false, nil); err != nil {
		return err
	}
	return nil
}

func configureQueue(ch *amqp.Channel, q Queue) (amqp.Queue, error) {
	log.Default().Println("Declaring queue " + q.Name)
	queue, err := ch.QueueDeclare(
		q.Name,
		q.Durable,
		false,       // delete when unused
		q.Exclusive, // exclusive
		false,       // noWait
		nil,         // arguments
	)

	if err != nil {
		return queue, err
	}
	e := q.Exchange
	if e.Name != "" {
		if err := declareExchange(ch, e); err != nil {
			return queue, err
		}

		return queue, ch.QueueBind(
			q.Name,
			q.FullPath(),
			q.Exchange.Name,
			false,
			nil,
		)
	}
	return queue, nil
}

func CreateSubscription(ch *amqp091.Channel, queue amqp.Queue, table amqp091.Table) (<-chan amqp091.Delivery, amqp091.Queue, error) {

	correlationId := table["CorrelationId"]
	fmt.Printf("Subscribing to %s with CorrelationID: %v\n", queue.Name, correlationId)

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		table,      // args
	)

	if err != nil {
		return nil, queue, err
	}

	return msgs, queue, err
}

func SubscribeToQueue(ch *amqp091.Channel, q Queue, table amqp091.Table) (<-chan amqp091.Delivery, amqp091.Queue, error) {
	queue, err := configureQueue(ch, q)
	if err != nil {
		return nil, queue, err
	}

	return CreateSubscription(ch, queue, table)
}

func (handler ConsumerHandler) start(ch *amqp091.Channel) error {
	go func() {

		msgs, _, err := SubscribeToQueue(ch, handler.queue, nil)
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
