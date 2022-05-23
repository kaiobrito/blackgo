package repository

import (
	"blackgo/deck"
	"blackgo/engine"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Deck deck.Deck

func (d *Deck) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := Deck{}
	err := json.Unmarshal(bytes, &result)
	*d = result
	return err
}

func (d Deck) Value() (driver.Value, error) {
	bytes, err := json.Marshal(d)
	return bytes, err
}

type GormBlackgo struct {
	gorm.Model

	ID         uuid.UUID  `gorm:"primaryKey"`
	UserDeck   Deck       `gorm:"type:jsonb`
	DealerDeck Deck       `gorm:"type:jsonb`
	Finished   *time.Time `gorm:"index"`
	Winner     engine.BlackGoWinner
	Stood      bool
}

func fromGame(game engine.Blackgo) GormBlackgo {
	return GormBlackgo{
		ID:         uuid.MustParse(game.ID),
		UserDeck:   Deck(game.UserDeck),
		DealerDeck: Deck(game.GetDealerDeck()),
		Winner:     game.Winner,
		Stood:      game.Stood,
		Finished:   nil,
	}
}

func (g GormBlackgo) toBlackgo() *engine.Blackgo {
	game := engine.CreateBlackgoWithDecks(deck.Deck(g.UserDeck), deck.Deck(g.DealerDeck))
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

func NewGormGameRepository() *GormGameRepository {
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
	repository.db.Save(&gorm_game)
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
