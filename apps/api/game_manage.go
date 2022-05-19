package main

import "blackgo/engine"

func CreateGame() engine.Blackgo {
	return engine.NewBlackgoGameWithShuffler(engine.DefaultShuffler())
}
