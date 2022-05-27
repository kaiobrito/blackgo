package messaging

import (
	"blackgo/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Manager struct {
	handlers      []ConsumerHandler
	relatedQueues []Queue
}

func NewManager(h []ConsumerHandler, q []Queue) *Manager {
	return &Manager{
		handlers:      h,
		relatedQueues: q,
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
	queues := m.relatedQueues
	hQueues := utils.Map(m.handlers, func(h ConsumerHandler) Queue { return h.queue })

	queues = append(queues, hQueues...)

	for _, q := range queues {
		if err := configureQueue(ch, q); err != nil {
			return err
		}
	}
	return nil
}

func configureQueue(ch *amqp.Channel, q Queue) error {
	e := q.Exchange

	if err := ch.ExchangeDeclare(e.Name, e.EType, e.Durable, false, false, false, nil); err != nil {
		return err
	}

	_, err := ch.QueueDeclare(
		q.Name,
		q.Durable,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	return ch.QueueBind(
		q.Name,
		q.FullPath(),
		q.Exchange.Name,
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
