package deck

type Card struct {
	Suit   CardSuit   `json:"suit" example:"spades"`
	Number CardNumber `json:"number" example:"A"`
}

func NewCard(suit CardSuit, number CardNumber) Card {
	return Card{suit, number}
}

func (card Card) ToString() string {
	return card.Number.Value() + " of " + card.Suit.Value()
}
