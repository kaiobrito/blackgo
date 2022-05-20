package engine

import (
	deck "blackgo/deck"
	cTypes "blackgo/deck/types"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartGame(t *testing.T) {
	game := NewBlackgoGame()
	game.Start()
	if game.Winner != NOONE {
		t.Errorf("There is winner yet")
	}

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

	if !reflect.DeepEqual(expectedDealer, game.dealerDeck) {
		t.Errorf("Dealer hand doesn't match")
	}
}

func TestJson(t *testing.T) {
	game := NewBlackgoGame()
	game.Start()
	assert.Equal(t, game.JSON(), map[string]any{
		"ID": game.ID,
		"user": map[string]any{
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
		},
		"dealer": map[string]any{
			"cards": []map[string]string{
				{
					"number": cTypes.C2.Value(),
					"suit":   cTypes.Spades.Value(),
				},
				{
					"number": cTypes.C3.Value(),
					"suit":   cTypes.Spades.Value(),
				},
			},
		},
		"winner": NOONE,
	})
}

func TestHit(t *testing.T) {
	game := NewBlackgoGame()
	game.Start()

	expected := deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C1),
		cTypes.NewCard(cTypes.Spades, cTypes.C4),
	}
	game.Hit()
	if !reflect.DeepEqual(expected, game.UserDeck) {
		t.Errorf("User hand doesn't match")
	}
	game.Hit()
	if game.Winner != USER {
		t.Errorf("User Has blackgo")
	}
}

func TestCheckWinner(t *testing.T) {
	game := NewBlackgoGame()
	game.Start()
	game.UserDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
	}
	game.checkWinner()
	if game.Winner != USER {
		t.Errorf("User won the game")
	}

	game.UserDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
	}
	game.checkWinner()

	if game.Winner != DEALER {
		t.Errorf("Dealer won the game. User is out of play")
	}

	game.UserDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
	}

	game.dealerDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
	}
	game.checkWinner()
	if game.Winner != NOONE {
		t.Errorf("The game isn't over")
	}

	game.UserDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
	}

	game.Stood = true
	game.checkWinner()
	if game.Winner != TIE {
		t.Errorf("The game was tied")
	}
}

func TestCheckWinnerAfterStand(t *testing.T) {
	game := NewBlackgoGame()
	game.Start()
	game.UserDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C5),
	}
	game.dealerDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
	}
	game.Stood = true
	game.checkWinner()
	if game.Winner != DEALER {
		t.Errorf("Dealer won. Higher score")
	}

	game.dealerDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.CA),
		cTypes.NewCard(cTypes.Spades, cTypes.C5),
	}
	game.UserDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
	}
	game.checkWinner()
	if game.Winner != USER {
		t.Errorf("User won. Higher score")
	}
}

func TestDealerOutOfPlay(t *testing.T) {
	game := NewBlackgoGame()
	game.Start()
	game.Stood = true
	game.dealerDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.C6),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
		cTypes.NewCard(cTypes.Spades, cTypes.C8),
	}
	game.UserDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.C4),
		cTypes.NewCard(cTypes.Spades, cTypes.CK),
		cTypes.NewCard(cTypes.Spades, cTypes.C6),
	}
	game.checkWinner()
	if game.Winner != USER {
		t.Errorf("User won. Dealer is out of play")
	}
}

func TestStand(t *testing.T) {
	game := NewBlackgoGame()
	game.Start()

	expected := deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.C2),
		cTypes.NewCard(cTypes.Spades, cTypes.C3),
		cTypes.NewCard(cTypes.Spades, cTypes.C4),
		cTypes.NewCard(cTypes.Spades, cTypes.C5),
		cTypes.NewCard(cTypes.Spades, cTypes.C6),
	}
	game.Stand()
	if !reflect.DeepEqual(expected, game.dealerDeck) {
		t.Errorf("Dealer deck not matching")
	}
}

func TestStandFromExistingValues(t *testing.T) {
	game := NewBlackgoGame()
	game.Start()

	expected := deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
	}
	game.dealerDeck = expected
	game.Stand()
	if !reflect.DeepEqual(expected, game.dealerDeck) {
		t.Errorf("Dealer deck not matching")
	}
}

func TestShuffler(t *testing.T) {
	called := false
	var shuffler IShuffler = func(d deck.Deck) deck.Deck {
		called = true
		return d
	}
	game := NewBlackgoGameWithShuffler(shuffler)
	game.Start()
	if !called {
		t.Errorf("Deck not shuffled")
	}
}
