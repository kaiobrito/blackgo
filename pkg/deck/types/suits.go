package deck

type CardSuit string

const (
	Spades   CardSuit = "spades"
	Diamonds CardSuit = "diamonds"
	Hearts   CardSuit = "hearts"
	Clubs    CardSuit = "clubs"
)

func AllCardSuits() []CardSuit {
	return []CardSuit{
		Spades, Diamonds, Hearts, Clubs,
	}
}
