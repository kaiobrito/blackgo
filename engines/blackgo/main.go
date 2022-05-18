package engine

import (
	"blackgo/deck"
)

type BlackGoWinner int64

const (
	NOONE BlackGoWinner = -1
	USER                = iota
	DEALER
	TIE
)

type Blackgo struct {
	d          deck.Deck
	UserDeck   deck.Deck
	DealerDeck deck.Deck
	Winner     BlackGoWinner
	Stood      bool
}

func (b *Blackgo) Start() {
	userDeck, newDeck := b.d.Deal(2)
	b.UserDeck = userDeck

	dealerDeck, newDeck := newDeck.Deal(2)
	b.DealerDeck = dealerDeck
	b.d = newDeck
	b.checkWinner()
}

func (b *Blackgo) checkWinner() {
	if highestValidCombination(b.UserDeck) == highestValidCombination(b.DealerDeck) && b.Stood {
		b.Winner = TIE
	} else if isOutOfPlay(b.UserDeck) {
		b.Winner = DEALER
	} else if checkBlackGo(b.UserDeck) {
		b.Winner = USER
	} else {
		b.Winner = NOONE
	}

}

func (b *Blackgo) Hit() {

}

func (b *Blackgo) Stand() {

}

func NewBlackgoGame() Blackgo {
	return Blackgo{
		d:          deck.GenerateDeck(),
		UserDeck:   nil,
		DealerDeck: nil,
		Winner:     NOONE,
		Stood:      false,
	}
}
