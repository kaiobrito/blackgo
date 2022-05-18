package engine

import (
	"blackgo/deck"
	"math/rand"
	"time"
)

type IShuffler func(deck.Deck) deck.Deck

func NoShuffler() IShuffler {
	return func(d deck.Deck) deck.Deck {
		return d
	}
}

func DefaultShuffler() IShuffler {
	return func(d deck.Deck) deck.Deck {
		rand.Seed(time.Now().UnixNano())

		rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
		return d
	}
}
