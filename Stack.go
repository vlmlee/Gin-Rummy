package main

// Stack - the discard pile in Gin Rummy.
type Stack []Card

// InitializeStack - puts the top card of the deck onto the stack.
func (d *Deck) InitializeStack() (stack Stack) {
	card := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	stack = append(stack, card)
	return
}

// DrawCard - picks up the top card of the stack.
func (s *Stack) DrawCard() (card Card) {
	card = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}
