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
	Exclusive  bool
}

func (q Queue) FullPath() string {
	return q.Name + "." + q.RoutingKey
}

type Message struct {
	Exchange      string
	CorrelationId string
	RoutingKey    string
	ReplyToName   *string
	Body          []byte
}
