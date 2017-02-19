package main

// StartNewGame initalizes the players, the deck, and deals cards to each
// player.
func StartNewGame() {
	RummyDeck := InitializeDeck()
	RummyStack := Stack{}
}

// Action describes what the player is going to do.
func Action() {

}

// Log will log the details of each turn, that is, what actions the player and
// the AI took.
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

}

func main() {
	// var score int
	StartNewGame()
	// Play again?
	// Tally score + replay with new scores
	// if ... {
	//		StartNewGame(score)
	// }
}
