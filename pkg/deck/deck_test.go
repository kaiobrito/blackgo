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

	if !reflect.DeepEqual(expected, myhand) {
		t.Errorf("My hand doesn't match")
	}
}

func TestJson(t *testing.T) {
	d := Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C1),
	}

	expected := map[string]any{
		"cards": []map[string]string{
			{
				"number": cTypes.CA.Value(),
				"suit":   cTypes.Spades.Value(),
			},
			{
				"number": cTypes.C1.Value(),
				"suit":   cTypes.Spades.Value(),
			},
		},
	}

	if !reflect.DeepEqual(expected, d.Json()) {
		t.Error("JSON doesn't match", expected, d.Json())
	}
}
