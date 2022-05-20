package controllers

import "blackgo/engine"

var Games map[string]*engine.Blackgo

func CreateGame() engine.Blackgo {
	return engine.NewBlackgoGameWithShuffler(engine.DefaultShuffler())
}

func init() {
	Games = map[string]*engine.Blackgo{}
}
