package main

// Stack is the discard pile in Gin Rummy.
type Stack []Card

// PickUp will take the card on top of the stack and put it into the player's // hand.
func (s *Stack) PickUp(hand *Hand) {
	hand.AddCardToHand(*s)
}
