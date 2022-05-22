package deck

import (
	"encoding/json"
	"fmt"
)

type Card struct {
	Number CardNumber `json:"number" example:"A"`
	Suit   CardSuit   `json:"suit" example:"spades"`
}

func (a *Card) UnmarshalJSON(b []byte) error {
	var result map[string]string

	if err := json.Unmarshal(b, &result); err != nil {
		fmt.Println(err)
		return err
	}
	*a = Card{
		Number: CardNumber(result["number"]),
		Suit:   CardSuit(result["suit"]),
	}
	return nil
}

func NewCard(suit CardSuit, number CardNumber) Card {
	return Card{number, suit}
}

func (card Card) ToString() string {
	return string(card.Number) + " of " + string(card.Suit)
}
