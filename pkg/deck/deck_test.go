package deck

import (
	cTypes "blackgo/deck/types"
	"reflect"
	"testing"
)

func TestDeal(t *testing.T) {
	deck := GenerateDeck()
	myhand, _ := deck.Deal(2)

	expected := Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C1),
	}

	myhand.Print()
	if !reflect.DeepEqual(expected, myhand) {
		t.Errorf("My hand doesn't match")
	}
}
