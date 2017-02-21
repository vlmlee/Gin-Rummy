package main

// Stack is the discard pile in Gin Rummy.
type Stack []Card

// InitalizeStack puts the top card of the deck onto the stack.
func (d *Deck) InitalizeStack() (stack Stack) {
	*d = (*d)[:len(*d)-1]
	card := (*d)[len(*d)-1]
	stack = append(stack, card)
	return
}

// DrawCard picks up the top card of the stack.
func (s *Stack) DrawCard() (card Card) {
	*s = (*s)[:len(*s)-1]
	return (*s)[len(*s)-1]
}
