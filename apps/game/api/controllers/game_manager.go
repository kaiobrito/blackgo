package controllers

import (
	"blackgo/engine"
)

var Games map[string]*engine.Blackgo

func CreateGame() engine.Blackgo {
	game := engine.NewBlackgoGameWithShuffler(engine.DefaultShuffler())
	Games[game.ID] = &game

	game.Start()

	return game
}

func init() {
	Games = map[string]*engine.Blackgo{}
}
