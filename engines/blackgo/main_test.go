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

	if !reflect.DeepEqual(expectedDealer, game.DealerDeck) {
		t.Errorf("Dealer hand doesn't match")
	}
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

	game.DealerDeck = deck.Deck{
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
	game.DealerDeck = deck.Deck{
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
		cTypes.NewCard(cTypes.Spades, cTypes.C10),
	}
	game.Stood = true
	game.checkWinner()
	if game.Winner != DEALER {
		t.Errorf("Dealer won. Higher score")
	}

	game.DealerDeck = deck.Deck{
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
