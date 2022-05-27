package queues

import "blackgo/messaging"

var GAMES_ACTION_EXCHANGE messaging.Exchange

var GAMES_CREATE_QUEUE messaging.Queue

func init() {
	GAMES_ACTION_EXCHANGE = messaging.Exchange{
		Name:    "games.direct",
		EType:   "direct",
		Durable: true,
	}

	GAMES_CREATE_QUEUE = messaging.Queue{
		Name:       "games.action",
		RoutingKey: "create",
		Exchange:   GAMES_ACTION_EXCHANGE,
		Durable:    true,
	}
}
