package main

import "testing"

func TestStackInitializes(t *testing.T) {
	deck := Deck{
		{13, "Clubs", "K"},
	}

	stack := deck.InitalizeStack()
	if len(stack) != 1 {
		t.Error("Stack did not initialize.")
	}
	return
}
