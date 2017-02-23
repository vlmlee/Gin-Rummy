package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartNewGame initalizes the players, the deck, and deals cards to each
// player.
func StartNewGame() (err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a name:")
	name, err := reader.ReadString('\n')

	if err != nil {
		return fmt.Errorf("an error occurred with the name")
	}

	p1 := &Player{name, Hand{}}
	p2 := &Player{"AI", Hand{}}
	RummyDeck := InitializeDeck()
	RummyStack := Stack{}

	RummyDeck.Deal(p1, p2)

	Action()
}

// Action describes what the player is going to do.
func Action() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("DRAW CARD --- PICKUP CARD --- CHECK MELDS --- CHECK POINTS")
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Unrecognized command.")
			continue
		}
		response = strings.ToUpper(strings.TrimSpace(response))
		switch response {
		case "DRAW":
			// Draw card from deck
		case "PICKUP":
			// Draw card from stack
		case "CHECK MELDS":
			// Check the melds in your hand
		case "CHECK TOTAL":
			// Check the total of points in your hand, values not melded
		}
	}

	fmt.Println("Discard a card from your hand.")

	// Pretty print hand

	// discard := reader.ReadString('\n')

	// DiscardCardFromHand(discard)

	// Pass turn to other player
}

// Log will log the details of each turn, that is, what actions the player and
// the AI took, top card on the stack.
func Log() {

}

// Turn keeps the state of which player's turn it is. They must either draw a card from the top of the discard pile or from the top of the deck. Then they must discard a card from their hand.
func Turn() {

}

// Knock ends the game and calculates the number of points the knocking player // has won or lost. AI and player will show hands here.
func Knock() {

}

// EndGame will immediately end the game and close the program.
func EndGame() {
	os.Exit(0)
}

func main() {
	var score int
	StartNewGame()
	// Play again?
	// Tally score + replay with new scores
	// if ... {
	//		StartNewGame(score)
	// }
}
