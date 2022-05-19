package controllers

import "blackgo/engine"

var games map[string]*engine.Blackgo

func CreateGame() engine.Blackgo {
	return engine.NewBlackgoGameWithShuffler(engine.DefaultShuffler())
}

func init() {
	games = map[string]*engine.Blackgo{}
}
