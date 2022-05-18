package engine

import (
	"blackgo/deck"
)

func NewBlackgoGame() Blackgo {
	return Blackgo{
		d:          deck.GenerateDeck(),
		UserDeck:   nil,
		DealerDeck: nil,
		Winner:     NOONE,
		Stood:      false,
		Shuffler:   NoShuffler(),
	}
}

func NewBlackgoGameWithShuffler(shuffler IShuffler) Blackgo {
	return Blackgo{
		d:          deck.GenerateDeck(),
		UserDeck:   nil,
		DealerDeck: nil,
		Winner:     NOONE,
		Stood:      false,
		Shuffler:   shuffler,
	}
}
