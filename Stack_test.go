package main

import "testing"

func TestStackInitializes(t *testing.T) {
	deck := Deck{
		{13, "Clubs", "K"},
	}

	stack := deck.InitializeStack()
	if len(stack) != 1 {
		t.Error("Stack did not initialize.")
	}
	return
}

func TestStackTopCard(t *testing.T) {
	deck := Deck{
		{13, "Clubs", "K"},
	}

	stack := deck.InitializeStack()
	if stack.PeekAtStack() != "KC" {
		t.Error("Stack did not return correct card.")
	}

	h := Hand{}
	h.DrawCard(&stack)

	if stack.PeekAtStack() != "No cards in the stack." {
		t.Error("Stack did not notify player that there are no cards in the stack.")
	}
}
