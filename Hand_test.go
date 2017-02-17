package main

import (
	"testing"
)

func TestHandSizeIsEqualToTen(t *testing.T) {
	hand := DealCards()
	if len(hand) != 10 {
		t.Error("Hand size is not 10!")
	}
	return
}

func TestPrettyPrintHand(t *testing.T) {
	input := []Card{
		{1, "Clubs", "A"}, {2, "Diamonds", "2"}, {3, "Spades", "3"},
		{4, "Hearts", "4"}, {5, "Clubs", "5"}, {6, "Diamonds", "6"},
		{7, "Hearts", "7"}, {11, "Clubs", "J"}, {12, "Spades", "Q"},
		{13, "Clubs", "K"},
	}
	pretty = input.PrettyPrint()
	if pretty != "AC 2D 3S 4H 5C 6D 7H JC QS KC" {
		t.Error("Pretty print did not prettify input!")
	}
	return
}
