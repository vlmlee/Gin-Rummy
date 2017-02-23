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
	RummyDeck.Deal(p1, p2)

	// While Knock is false, keep the players in a loop that handle turns.
	// We can set v := false
	for v := Knock(); v; {
		if turn == p1 {
			Action(p1, &RummyDeck, &RummyStack, &v)
			turn = p2
		} else {
			Action(p2, &RummyDeck, &RummyStack, &v)
			turn = p1
		}
	}

	fmt.Printf("Player score: %d AI score: %d \n Play again? (Y/N)", pScore, AIScore)
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
func Action(p *Player, deck *Deck, stack *Stack, knock *bool) {
	reader := bufio.NewReader(os.Stdin)

ACTIONS:
	for {
		fmt.Printf("Card on stack: %s \nYour hand: %s \n", stack.PeekAtStack(), p.PrettyPrintHand())
		fmt.Printf("\nWhat would you like to do, %s?\n 1. DRAW CARD\n 2. PICKUP CARD FROM STACK\n 3. CHECK MELDS\n 4. CHECK POINTS\n", p.name)
		response, err := reader.ReadString('\n')
		response = strings.TrimRight(response, "\n")
		if err != nil {
			fmt.Println("Unrecognized command.")
			continue
		}
		response = strings.ToUpper(strings.TrimSpace(response))
		switch response {
		case "1", "DRAW CARD":
			p.Hand.DrawCard(deck)
			break ACTIONS
		case "2", "PICKUP CARD FROM STACK":
			p.Hand.DrawCard(stack)
			break ACTIONS
		case "3", "CHECK MELDS":
			melds := p.Hand.CheckMelds()
			fmt.Printf("\n%s\n", melds.PrettyPrintMelds())
		case "4", "CHECK POINTS":
			// Check the total of points in your hand, values not melded
		default:
			fmt.Printf("\nUnrecognized command. Try again.\n\n")
		}
	}

	fmt.Println("Now, you must discard a card from your hand.")

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
	fmt.Println("Enter your name:")
	name, err := reader.ReadString('\n')
	name = strings.TrimRight(name, "\n")
	fmt.Printf("\n")

	if err != nil {
		fmt.Printf("An error occurred with the name")
		os.Exit(0)
	}

	fmt.Printf("Welcome %s! Lets play a game of Gin Rummy!\n", name)
	fmt.Printf("\n")
	StartNewGame(&name, &pScore, &AIScore)
}
