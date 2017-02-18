package main

// Player has the actions AddCardToHand, DiscardCardFromHand, Knock, and EndGame
type Player interface {
	AddCardToHand()
	DiscardCardFromHand()
	Knock()
	EndGame()
}
