package engine

import (
	"blackgo/deck"
	dTypes "blackgo/deck/types"
	"testing"
)

func TestCheckBlackgoWithLetters(t *testing.T) {
	hand10 := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CA),
		dTypes.NewCard(dTypes.Clubs, dTypes.C10),
	}
	handJ := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CJ),
		dTypes.NewCard(dTypes.Clubs, dTypes.CA),
	}
	handQ := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CQ),
		dTypes.NewCard(dTypes.Clubs, dTypes.CA),
	}
	handK := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CK),
		dTypes.NewCard(dTypes.Clubs, dTypes.CA),
	}

	if !checkBlackGo(hand10) {
		t.Errorf("A + 10 IS a blackgo")
	}
	if !checkBlackGo(handJ) {
		t.Errorf("A + J IS a blackgo")
	}
	if !checkBlackGo(handQ) {
		t.Errorf("A + Q IS a blackgo")
	}
	if !checkBlackGo(handK) {
		t.Errorf("A + K IS a blackgo")
	}
}

func TestCheckBlackgoJQK(t *testing.T) {
	handJ := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CJ),
		dTypes.NewCard(dTypes.Clubs, dTypes.C10),
	}
	handQ := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CQ),
		dTypes.NewCard(dTypes.Clubs, dTypes.C9),
	}
	handK := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CK),
		dTypes.NewCard(dTypes.Clubs, dTypes.C8),
	}

	if checkBlackGo(handJ) {
		t.Errorf("J+10 isn't a blackgo")
	}
	if checkBlackGo(handQ) {
		t.Errorf("Q+9 isn't a blackgo")
	}
	if checkBlackGo(handK) {
		t.Errorf("K+8 isn't a blackgo")
	}
}

func TestMultipleAces(t *testing.T) {
	hand2 := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CA),
		dTypes.NewCard(dTypes.Diamonds, dTypes.CA),
		dTypes.NewCard(dTypes.Clubs, dTypes.C9),
	}

	if !checkBlackGo(hand2) {
		t.Errorf("AA+9 IS a blackgo")
	}

	hand3 := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CA),
		dTypes.NewCard(dTypes.Diamonds, dTypes.CA),
		dTypes.NewCard(dTypes.Hearts, dTypes.CA),
		dTypes.NewCard(dTypes.Clubs, dTypes.C8),
	}

	if !checkBlackGo(hand3) {
		t.Errorf("AAA+8 IS a blackgo")
	}

	hand4 := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.CA),
		dTypes.NewCard(dTypes.Diamonds, dTypes.CA),
		dTypes.NewCard(dTypes.Hearts, dTypes.CA),
		dTypes.NewCard(dTypes.Spades, dTypes.CA),
		dTypes.NewCard(dTypes.Clubs, dTypes.C7),
	}

	if !checkBlackGo(hand4) {
		t.Errorf("AAA+7 IS a blackgo")
	}
}

func TestNumbersOnly(t *testing.T) {
	hand := deck.Deck{
		dTypes.NewCard(dTypes.Clubs, dTypes.C1),
		dTypes.NewCard(dTypes.Diamonds, dTypes.C2),
		dTypes.NewCard(dTypes.Diamonds, dTypes.C3),
		dTypes.NewCard(dTypes.Clubs, dTypes.C4),
		dTypes.NewCard(dTypes.Clubs, dTypes.C5),
		dTypes.NewCard(dTypes.Clubs, dTypes.C6),
	}

	if !checkBlackGo(hand) {
		t.Errorf("1+2+3+4+5+6 IS a blackgo")
	}

}
