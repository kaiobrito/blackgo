package deck

import (
	cTypes "blackgo/deck/types"
	"blackgo/utils"
	"fmt"
)

// The Deck type refers to a slice of strings
type Deck []cTypes.Card

func (deck Deck) Print() {
	fmt.Println(deck.ToString())
}

func (deck Deck) ToString() string {
	return utils.Reduce(deck, func(o string, t cTypes.Card) string {
		return o + t.ToString() + "\n"
	}, "")
}

func (deck Deck) Deal(handsize int) (Deck, Deck) {
	return deck[:handsize], deck[handsize:]
}

func GenerateDeck() Deck {
	suits := cTypes.AllCardSuits()
	numbers := cTypes.AllCardNumbers()

	deck := Deck{}

	for _, suit := range suits {
		for _, number := range numbers {
			deck = append(deck, cTypes.NewCard(suit, number))
		}
	}
	return deck
}
