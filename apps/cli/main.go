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
	game = engine.NewBlackgoGameWithShuffler(engine.DefaultShuffler())

	reader := bufio.NewReader(os.Stdin)

	game.Start()
	fmt.Println("New game has started!")
	fmt.Print("---------------------\n\n\n")

	for {

		fmt.Println("Your cards")

		for _, card := range game.UserDeck {
			fmt.Println(card.ToString())
		}
		fmt.Println("---------------------")
		fmt.Println("Dealer's cards")
		fmt.Println(game.DealerDeckAsString())

		fmt.Print("\n---------------------\n\n\n")
		fmt.Println("Type:")
		fmt.Println("1: To hit")
		fmt.Println("2: To stand")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "1" {
			game.Hit()
		}
		if text == "2" {
			game.Stand()
		}

		if game.Winner != engine.NOONE {
			fmt.Print("\n---------------------\n\n\n")
			fmt.Println("The winner is: " + game.Winner.ToString())
			fmt.Println("User deck: ")
			game.UserDeck.Print()

			fmt.Println("Dealer deck: ")
			fmt.Println(game.DealerDeckAsString())
			break
		}
	}
}
