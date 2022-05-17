package cards

import (
	"fmt"

	cTypes "blackgo/deck/types"
)

// The Deck type refers to a slice of strings
type Deck []cTypes.Card

func (deck Deck) Print() {
	for _, card := range deck {
		fmt.Println(card.ToString())
	}
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
