package repository

import (
	"blackgo/engine"
)

var games map[string]*engine.Blackgo

func CreateGame() *engine.Blackgo {
	game := engine.NewBlackgoGameWithShuffler(engine.DefaultShuffler())
	game.Start()
	SaveGame(&game)

	return &game
}

func SaveGame(game *engine.Blackgo) {
	games[game.ID] = game
}

func GetGameById(id string) *engine.Blackgo {
	return games[id]
}

func GetAllGames() map[string]*engine.Blackgo {
	return games
}

func DeleteAll() {
	games = map[string]*engine.Blackgo{}
}

func init() {
	DeleteAll()
}