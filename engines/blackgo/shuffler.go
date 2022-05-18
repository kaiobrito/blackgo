package engine

import "blackgo/deck"

type IShuffler func(deck.Deck) deck.Deck

func NoShuffler() IShuffler {
	return func(d deck.Deck) deck.Deck {
		return d
	}
}
