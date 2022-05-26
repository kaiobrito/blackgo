package messaging

import (
	"blackgo/utils"

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
	queues := utils.Map(m.handlers, func(h ConsumerHandler) string { return h.queue })

	for _, queue := range queues {
		_, err := ch.QueueDeclare(
			queue,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return nil
		}
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
