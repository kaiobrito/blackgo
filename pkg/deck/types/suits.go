package deck

import "encoding/json"

type CardSuit int64

func (a CardSuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Value())
}

const (
	Spades CardSuit = iota
	Diamonds
	Hearts
	Clubs
)

func AllCardSuits() []CardSuit {
	return []CardSuit{
		Spades, Diamonds, Hearts, Clubs,
	}
}

func (suit CardSuit) Value() string {
	switch suit {
	case Spades:
		return "spades"
	case Diamonds:
		return "diamonds"
	case Hearts:
		return "hearts"
	case Clubs:
		return "clubs"
	}
	return "unknown"
}
