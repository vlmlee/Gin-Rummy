package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartNewGame initalizes the players, the deck, and deals cards to each
// player.
func StartNewGame(name *string, pScore, AIScore *int) (err error) {
	p1 := &Player{*name, Hand{}}
	p2 := &Player{"AI", Hand{}}
	turn := p1
	RummyDeck := InitializeDeck()
	RummyStack := RummyDeck.InitializeStack()
	fmt.Println(RummyStack)
	RummyDeck.Deal(p1, p2)

	// While Knock is false, keep the players in a loop that handle turns.
	// We can set v := false
	for v := Knock(); v; {
		if turn == p1 {
			Action(p1, &v)
			turn = p2
		} else {
			Action(p2, &v)
			turn = p1
		}
	}

	fmt.Printf("Player score: %d AI score: %d \n", pScore, AIScore)
	fmt.Println("Play again? (Y/N)")
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	response = strings.ToUpper(strings.TrimSpace(response))

	// Start a new game
	if response == "Y" {
		StartNewGame(name, pScore, AIScore)
	}
	return
}

// Action describes what the player is going to do.
func Action(p *Player, knock *bool) {
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
func Log(p *Player, action string) {

}

// Knock ends the game and calculates the number of points the knocking player // has won or lost. AI and player will show hands here.
func Knock() bool {
	return true
}

func main() {
	pScore := 0
	AIScore := 0

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a name:")
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("An error occurred with the name")
		os.Exit(0)
	}

	StartNewGame(&name, &pScore, &AIScore)
}
