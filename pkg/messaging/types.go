package messaging

type Exchange struct {
	Name    string
	EType   string
	Durable bool
}

type Queue struct {
	Name       string
	Durable    bool
	RoutingKey string
	Exchange   Exchange
}

func (q Queue) FullPath() string {
	return q.Name + "." + q.RoutingKey
}

type Message struct {
	CorrelationId string
	Queue         Queue
	ReplyTo       *Queue
	Body          []byte
}
