package repository

import (
	"blackgo/engine"
)

type IGameRepository interface {
	CreateGame() *engine.Blackgo
	SaveGame(game *engine.Blackgo)
	GetGameById(id string) *engine.Blackgo
	GetAllGames() map[string]*engine.Blackgo
	DeleteAll()
}

var repository IGameRepository

func CreateGame() *engine.Blackgo {
	return repository.CreateGame()
}

func SaveGame(game *engine.Blackgo) {
	repository.SaveGame(game)
}

func GetGameById(id string) *engine.Blackgo {
	return repository.GetGameById(id)
}

func GetAllGames() map[string]*engine.Blackgo {
	return repository.GetAllGames()
}

func DeleteAll() {
	repository.DeleteAll()
}

func init() {
	repository = NewInMemoryRepository()
	//repository = NewGormGameRepository()
}
