package repository

import (
	"blackgo/deck"
	"blackgo/engine"
	"encoding/json"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormBlackgo struct {
	gorm.Model

	ID         uuid.UUID `gorm:"primaryKey"`
	UserDeck   datatypes.JSON
	DealerDeck datatypes.JSON
	Finished   *time.Time `gorm:"index"`
	Winner     engine.BlackGoWinner
	Stood      bool
}

func fromGame(game engine.Blackgo) GormBlackgo {
	ud, _ := json.Marshal(game.UserDeck)
	dd, _ := json.Marshal(game.DealerDeck)

	return GormBlackgo{
		ID:         uuid.MustParse(game.ID),
		UserDeck:   datatypes.JSON(ud),
		DealerDeck: datatypes.JSON(dd),
		Winner:     game.Winner,
		Stood:      game.Stood,
		Finished:   nil,
	}
}

func (g GormBlackgo) toBlackgo() *engine.Blackgo {
	var ud deck.Deck
	json.Unmarshal([]byte(g.UserDeck.String()), &ud)
	var dd deck.Deck
	json.Unmarshal([]byte(g.UserDeck.String()), &dd)

	game := engine.CreateBlackgoWithDecks(ud, dd)
	game.ID = g.ID.String()
	game.Winner = g.Winner
	game.Stood = g.Stood
	game.Shuffler = engine.DefaultShuffler()
	game.Shuffle()

	return game
}

type GormGameRepository struct {
	db *gorm.DB
}

func NewGormGameRepository() IGameRepository {
	dns := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&GormBlackgo{})
	return &GormGameRepository{db: db}
}

func (repository GormGameRepository) CreateGame() *engine.Blackgo {
	game := engine.NewBlackgoGameWithShuffler(engine.DefaultShuffler())
	game.Start()
	repository.SaveGame(&game)
	return &game
}

func (repository GormGameRepository) SaveGame(game *engine.Blackgo) {
	gorm_game := fromGame(*game)
	if err := repository.db.Save(&gorm_game).Error; err != nil {
		game = nil
	}
}

func (repository GormGameRepository) GetGameById(id string) *engine.Blackgo {
	var game GormBlackgo
	repository.db.First(&game, "id", id)

	return game.toBlackgo()
}

func (repository GormGameRepository) GetAllGames() map[string]*engine.Blackgo {
	return nil
}

func (repository GormGameRepository) DeleteAll() {

}
