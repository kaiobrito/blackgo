package engine

import (
	"blackgo/deck"
	dTypes "blackgo/deck/types"
	"blackgo/engine/exceptions"
	"blackgo/utils"
)

type BlackGoWinner int64

const (
	NOONE BlackGoWinner = -1
	USER                = iota
	DEALER
	TIE
)

func (winner BlackGoWinner) ToString() string {
	switch winner {
	case USER:
		return "User"
	case DEALER:
		return "Dealer"
	case TIE:
		return "Tie"
	default:
		return "Game still running"
	}
}

type Blackgo struct {
	ID         string        `json:"ID" example:"foobar"`
	d          deck.Deck     `json:"-"`
	UserDeck   deck.Deck     `json:"user" `
	DealerDeck deck.Deck     `json:"dealer" `
	Winner     BlackGoWinner `json:"winner" `
	Stood      bool          `json:"stood" `
	Shuffler   IShuffler     `json:"-"`
}

func CreateBlackgoWithDecks(uDeck deck.Deck, DeDeck deck.Deck) *Blackgo {
	remainingDeck := utils.Filter(deck.GenerateDeck(), func(card dTypes.Card) bool {
		return uDeck.Contains(card) || DeDeck.Contains(card)
	})

	return &Blackgo{
		DealerDeck: DeDeck,
		UserDeck:   uDeck,
		d:          remainingDeck,
	}
}

func (b *Blackgo) GetDealerDeck() deck.Deck {
	return b.DealerDeck
}

func (b *Blackgo) Start() {
	b.Shuffle()
	userDeck, newDeck := b.d.Deal(2)
	b.UserDeck = userDeck

	dealerDeck, newDeck := newDeck.Deal(2)
	b.DealerDeck = dealerDeck
	b.d = newDeck
	b.checkWinner()
}

func (b *Blackgo) Shuffle() {
	b.d = b.Shuffler(b.d)
}

func (b *Blackgo) checkWinner() {
	if highestValidCombination(b.UserDeck) == highestValidCombination(b.DealerDeck) && b.Stood {
		b.Winner = TIE
	} else if isOutOfPlay(b.UserDeck) {
		b.Winner = DEALER
	} else if isOutOfPlay(b.DealerDeck) && b.Stood {
		b.Winner = USER
	} else if checkBlackGo(b.UserDeck) {
		b.Winner = USER
	} else if highestValidCombination(b.UserDeck) > highestValidCombination(b.DealerDeck) && b.Stood {
		b.Winner = USER
	} else if highestValidCombination(b.UserDeck) < highestValidCombination(b.DealerDeck) && b.Stood {
		b.Winner = DEALER
	} else {
		b.Winner = NOONE
	}

}

func (b *Blackgo) Hit() error {
	if b.Winner != NOONE {
		return exceptions.ErrGameIsOver
	}
	newCard, newDeck := b.d.Deal(1)
	b.UserDeck = append(b.UserDeck, newCard...)
	b.d = newDeck
	b.checkWinner()
	return nil
}

func (b *Blackgo) Stand() {
	b.Stood = true
	if highestValidCombination(b.DealerDeck) < 17 {
		for {
			newCard, newDeck := b.d.Deal(1)
			b.DealerDeck = append(b.DealerDeck, newCard...)
			b.d = newDeck

			if highestValidCombination(b.DealerDeck) >= 17 {
				break
			}
		}
	}
	b.checkWinner()
}

func (b Blackgo) DealerDeckAsString() string {
	if b.Stood {
		return b.DealerDeck.ToString()
	}
	return b.DealerDeck[0].ToString()
}
