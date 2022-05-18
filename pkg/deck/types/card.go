package deck

type Card struct {
	Suit   CardSuit
	Number CardNumber
}

func NewCard(suit CardSuit, number CardNumber) Card {
	return Card{suit, number}
}

func (card Card) ToString() string {
	return card.Number.Value() + " of " + card.Suit.Value()
}
