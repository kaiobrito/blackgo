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
	if err := m.startConsumers(ch); err != nil {
		return err
	}

	return nil
}

func (m Manager) startConsumers(ch *amqp.Channel) error {
	for _, handler := range m.handlers {
		if err := handler.start(ch); err != nil {
			return err
		}
	}
	return nil
}
