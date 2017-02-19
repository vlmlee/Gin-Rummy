package main

import "testing"

func TestMaxHandSizeIsEqualToEleven(t *testing.T) {
	deck := InitializeDeck()
	hand := &Hand{}

	for i := 0; i < 11; i++ {
		deck, _ = deck.DrawCard(hand)
	}

	if len(*hand) != 11 {
		t.Error("Hand size is not 11!")
	}

	deck, err := deck.DrawCard(hand)

	if err == nil {
		t.Error("Did not throw error when hand size became more than 11!")
	}
	return
}

func TestPrettyPrintHand(t *testing.T) {
	hand := Hand{
		{1, "Clubs", "A"}, {2, "Diamonds", "2"}, {3, "Spades", "3"},
		{4, "Hearts", "4"}, {5, "Clubs", "5"}, {6, "Diamonds", "6"},
		{7, "Hearts", "7"}, {11, "Clubs", "J"}, {12, "Spades", "Q"},
		{13, "Clubs", "K"},
	}

	prettyHand := hand.PrettyPrint()

	if prettyHand != "AC 2D 3S 4H 5C 6D 7H JC QS KC" {
		t.Error("Pretty print did not prettify input!")
	}

	return
}
