package engine

import "blackgo/deck"

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
	d          deck.Deck
	UserDeck   deck.Deck
	dealerDeck deck.Deck
	Winner     BlackGoWinner
	Stood      bool
	Shuffler   IShuffler
}

func (b *Blackgo) Start() {
	b.d = b.Shuffler(b.d)
	userDeck, newDeck := b.d.Deal(2)
	b.UserDeck = userDeck

	dealerDeck, newDeck := newDeck.Deal(2)
	b.dealerDeck = dealerDeck
	b.d = newDeck
	b.checkWinner()
}

func (b *Blackgo) checkWinner() {
	if highestValidCombination(b.UserDeck) == highestValidCombination(b.dealerDeck) && b.Stood {
		b.Winner = TIE
	} else if isOutOfPlay(b.UserDeck) {
		b.Winner = DEALER
	} else if isOutOfPlay(b.dealerDeck) && b.Stood {
		b.Winner = USER
	} else if checkBlackGo(b.UserDeck) {
		b.Winner = USER
	} else if highestValidCombination(b.UserDeck) > highestValidCombination(b.dealerDeck) && b.Stood {
		b.Winner = USER
	} else if highestValidCombination(b.UserDeck) < highestValidCombination(b.dealerDeck) && b.Stood {
		b.Winner = DEALER
	} else {
		b.Winner = NOONE
	}

}

func (b *Blackgo) Hit() {
	newCard, newDeck := b.d.Deal(1)
	b.UserDeck = append(b.UserDeck, newCard...)
	b.d = newDeck
	b.checkWinner()
}

func (b *Blackgo) Stand() {
	b.Stood = true
	if highestValidCombination(b.dealerDeck) < 17 {
		for {
			newCard, newDeck := b.d.Deal(1)
			b.dealerDeck = append(b.dealerDeck, newCard...)
			b.d = newDeck

			if highestValidCombination(b.dealerDeck) >= 17 {
				break
			}
		}
	}
	b.checkWinner()
}

func (b Blackgo) DealerDeckAsString() string {
	if b.Stood {
		return b.dealerDeck.ToString()
	}
	return b.dealerDeck[0].ToString()
}
