package main

import (
	"blackgo/engine"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var game engine.Blackgo

func main() {

	fmt.Println("Welcome to Blackgo!")
	game = engine.NewBlackgoGame()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---------------------")
	fmt.Println("Type:")
	fmt.Println("1: To start a new Game")
	fmt.Println("2: Finish the came")

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if text != "1" {
		return
	}

	game.Start()
	fmt.Println("New game has started!")
	fmt.Println("---------------------")

	for {

		fmt.Println("Your cards")

		for _, card := range game.UserDeck {
			fmt.Println(card.ToString())
		}
		fmt.Println("---------------------")
		fmt.Println("Dealer's cards")
		fmt.Println(game.DealerDeck[0].ToString())
		fmt.Println("* of *")

		fmt.Print("\n---------------------\n\n\n")
		fmt.Println("Type:")
		fmt.Println("1: To hit")
		fmt.Println("2: To stand")

		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "1" {
			game.Hit()
		}
		if text == "2" {
			game.Stand()
		}

		if game.Winner != engine.NOONE {
			fmt.Println("The winner is: " + game.Winner.ToString())
			break
		}
	}
}