package messaging

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Manager struct {
	handlers []ConsumerHandler
}

func NewManager(h []ConsumerHandler) *Manager {
	return &Manager{
		handlers: h,
	}
}

func (m Manager) Start(ch *amqp.Channel) error {
	if err := m.declareQueues(ch); err != nil {
		return err
	}
	if err := m.startConsumers(ch); err != nil {
		return err
	}

	return nil
}

func (m Manager) declareQueues(ch *amqp.Channel) error {
	for _, handler := range m.handlers {
		if err := configureConsumer(ch, handler); err != nil {
			return nil
		}
	}
	return nil
}

func configureConsumer(ch *amqp.Channel, handler ConsumerHandler) error {
	e := handler.queue.Exchange

	if err := ch.ExchangeDeclare(e.Name, e.EType, e.Durable, false, false, false, nil); err != nil {
		return err
	}

	_, err := ch.QueueDeclare(
		handler.queue.Name,
		handler.queue.Durable,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	return ch.QueueBind(
		handler.queue.Name,
		handler.queue.FullPath(),
		handler.queue.Exchange.Name,
		false,
		nil,
	)
}

func (m Manager) startConsumers(ch *amqp.Channel) error {
	for _, handler := range m.handlers {
		if err := handler.start(ch); err != nil {
			return err
		}
	}
	return nil
}
