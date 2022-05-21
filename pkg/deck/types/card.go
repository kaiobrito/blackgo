package deck

type Card struct {
	Number CardNumber `json:"number" example:"A"`
	Suit   CardSuit   `json:"suit" example:"spades"`
}

func NewCard(suit CardSuit, number CardNumber) Card {
	return Card{number, suit}
}

func (card Card) ToString() string {
	return card.Number.Value() + " of " + card.Suit.Value()
}
