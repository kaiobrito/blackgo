package repository

import (
	"blackgo/engine"
	"sync"
)

type InMemoryGameRepository struct {
	mux   sync.Mutex
	games map[string]*engine.Blackgo
}

func NewInMemoryRepository() *InMemoryGameRepository {
	return &InMemoryGameRepository{games: make(map[string]*engine.Blackgo)}
}

func (repository *InMemoryGameRepository) CreateGame() *engine.Blackgo {
	game := engine.NewBlackgoGameWithShuffler(engine.DefaultShuffler())
	game.Start()
	repository.SaveGame(&game)

	return &game
}

func (repository *InMemoryGameRepository) SaveGame(game *engine.Blackgo) {
	repository.mux.Lock()
	repository.games[game.ID] = game
	repository.mux.Unlock()
}

func (repository *InMemoryGameRepository) GetGameById(id string) *engine.Blackgo {
	repository.mux.Lock()
	game := repository.games[id]
	repository.mux.Unlock()
	return game
}

func (repository *InMemoryGameRepository) GetAllGames() map[string]*engine.Blackgo {
	return repository.games
}

func (repository *InMemoryGameRepository) DeleteAll() {
	repository.mux.Lock()
	repository.games = map[string]*engine.Blackgo{}
	repository.mux.Unlock()
}
