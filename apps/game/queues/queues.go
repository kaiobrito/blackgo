package queues

import "blackgo/messaging"

// Exchanges
var GAMES_ACTION_EXCHANGE messaging.Exchange
var GAMES_QUERY_EXCHANGE messaging.Exchange

// Queues
var GAMES_CREATE_QUEUE messaging.Queue
var GAMES_QUERY_QUEUE messaging.Queue

func init() {
	configExchanges()

	GAMES_CREATE_QUEUE = messaging.Queue{
		Name:       "games.action",
		RoutingKey: "create",
		Exchange:   GAMES_ACTION_EXCHANGE,
		Durable:    true,
	}

	GAMES_QUERY_QUEUE = messaging.Queue{
		Name:       "games.query",
		RoutingKey: "get_by_id",
		Exchange:   GAMES_QUERY_EXCHANGE,
		Durable:    true,
	}
}

func configExchanges() {
	GAMES_ACTION_EXCHANGE = messaging.Exchange{
		Name:    "games.direct",
		EType:   "direct",
		Durable: true,
	}
	GAMES_QUERY_EXCHANGE = messaging.Exchange{
		Name:    "games.query",
		EType:   "direct",
		Durable: false,
	}
}
