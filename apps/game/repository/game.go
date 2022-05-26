package repository

import (
	"blackgo/engine"
	"log"
	"os"
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

type RepositoryConstructor func() IGameRepository

func init() {

	REPOSITORIES := map[string]RepositoryConstructor{
		"GORM":   NewGormGameRepository,
		"AMQP":   NewAMQPGameRepository,
		"MEMORY": NewInMemoryRepository,
	}
	rType := os.Getenv("REPOSITORY_TYPE")
	constructor := REPOSITORIES[rType]
	if constructor != nil {
		log.Println("Starting with repository " + rType)
		repository = constructor()
	} else {
		log.Println("fallback: Using in-memory repository")
		repository = NewInMemoryRepository()
	}
}
