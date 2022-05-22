package repository

import "blackgo/engine"

type InMemoryGameRepository struct {
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
	repository.games[game.ID] = game
}

func (repository *InMemoryGameRepository) GetGameById(id string) *engine.Blackgo {
	return repository.games[id]
}

func (repository *InMemoryGameRepository) GetAllGames() map[string]*engine.Blackgo {
	return repository.games
}

func (repository *InMemoryGameRepository) DeleteAll() {
	repository.games = map[string]*engine.Blackgo{}
}
