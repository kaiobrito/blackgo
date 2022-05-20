package engine

import (
	"blackgo/deck"

	"github.com/google/uuid"
)

func NewBlackgoGame() Blackgo {
	return Blackgo{
		ID:         uuid.NewString(),
		d:          deck.GenerateDeck(),
		UserDeck:   nil,
		dealerDeck: nil,
		Winner:     NOONE,
		Stood:      false,
		Shuffler:   NoShuffler(),
	}
}

func NewBlackgoGameWithShuffler(shuffler IShuffler) Blackgo {
	return Blackgo{
		ID:         uuid.NewString(),
		d:          deck.GenerateDeck(),
		UserDeck:   nil,
		dealerDeck: nil,
		Winner:     NOONE,
		Stood:      false,
		Shuffler:   shuffler,
	}
}
