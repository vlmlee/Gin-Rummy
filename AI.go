package main

// AI can cheat and peek at all cards in the stack.

// The general strategy is to play greedily. The AI will ignore the player's
// actions and try to as quickly as possible to knock and end the game. This
// means it will simply look for nearest neighbors of cards in its hand that
// is currently not in a meld. If the card on the stack is not within 2 range,
// it will draw a card from the deck. If it has to decide between two cards
// with one pairing to discard, it will always discard the greater value card.

// IsNearestNeighbor will check if card on the top of the stack is a nearest
// neighbor.
func IsNearestNeighbor() {

}

// ChooseCardToDiscard will decide between cards to discard.
func ChooseCardToDiscard() {

}
