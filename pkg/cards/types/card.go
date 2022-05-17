package cards

type Card struct {
	suit   CardSuit
	number CardNumber
}

func NewCard(suit CardSuit, number CardNumber) Card {
	return Card{suit, number}
}

func (card Card) ToString() string {
	return card.number.Value() + " of " + card.suit.Value()
}
