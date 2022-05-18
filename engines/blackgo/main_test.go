package engine

import (
	deck "blackgo/deck"
	cTypes "blackgo/deck/types"
	"reflect"
	"testing"
)

func TestStartGame(t *testing.T) {
	game := NewBlackgoGame()
	game.Start()

	expected := deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C1),
	}

	if !reflect.DeepEqual(expected, game.UserDeck) {
		t.Errorf("User hand doesn't match")
	}

	expectedDealer := deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.C2),
		cTypes.NewCard(cTypes.Spades, cTypes.C3),
	}

	if !reflect.DeepEqual(expectedDealer, game.DealerDeck) {
		t.Errorf("Dealer hand doesn't match")
	}
}
