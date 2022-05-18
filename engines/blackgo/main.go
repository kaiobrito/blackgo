package engine

import (
	"blackgo/deck"
)

type BlackGoWinner int64

const (
	USER BlackGoWinner = iota
	DEALER
)

type Blackgo struct {
	d          deck.Deck
	UserDeck   deck.Deck
	DealerDeck deck.Deck
	Winner     *BlackGoWinner
}

func (b *Blackgo) Start() {
	userDeck, newDeck := b.d.Deal(2)
	b.UserDeck = userDeck

	dealerDeck, newDeck := newDeck.Deal(2)
	b.DealerDeck = dealerDeck
	b.d = newDeck
}

func (b *Blackgo) checkWinner() {

}

func (b *Blackgo) Hit() {

}

func (b *Blackgo) Stand() {

}

func NewBlackgoGame() Blackgo {
	return Blackgo{d: deck.GenerateDeck()}
}
