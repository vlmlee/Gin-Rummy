package main

import (
	"reflect"
	"testing"
)

func TestMaxHandSizeIsEqualToEleven(t *testing.T) {
	deck := InitializeDeck()
	hand := &Hand{}
	for i := 0; i < 11; i++ {
		hand.DrawCard(&deck)
	}
	if len(*hand) != 11 {
		t.Error("Hand size is not 11.")
	}

	_, err := hand.DrawCard(&deck)
	if err == nil {
		t.Error("Did not throw error when hand size became more than 11.")
	}
	return
}

func TestPrettyPrintHand(t *testing.T) {
	hand := Hand{
		{13, "Clubs", "K"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {6, "Diamonds", "6"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {5, "Clubs", "5"}, {7, "Hearts", "7"},
		{11, "Clubs", "J"},
	}

	prettyHand := hand.PrettyPrintHand()
	if prettyHand != "AC 2D 3S 4H 5C 6D 7H JC QS KC" {
		t.Error("Pretty print did not prettify input.")
	}
	return
}

func TestDiscardCard(t *testing.T) {
	hand := Hand{
		{13, "Clubs", "K"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {6, "Diamonds", "6"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {5, "Clubs", "5"}, {7, "Hearts", "7"},
		{11, "Clubs", "J"},
	}

	stack := &Stack{}
	hand.DiscardCard(Card{4, "Hearts", "4"}, stack)

	if reflect.DeepEqual(hand, []Hand{{
		{13, "Clubs", "K"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{6, "Diamonds", "6"}, {12, "Spades", "Q"}, {3, "Spades", "3"},
		{5, "Clubs", "5"}, {7, "Hearts", "7"}, {11, "Clubs", "J"}},
	}) {
		t.Error("Hand did not discard card 4 of Hearts.")
	}

	if reflect.DeepEqual(stack, Stack{Card{4, "Hearts", "4"}}) {
		t.Error("Stack did not put card on top of pile.")
	}
	return
}

func TestCheckTotal(t *testing.T) {
	hand := Hand{
		{2, "Clubs", "2"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {3, "Diamonds", "3"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {4, "Diamonds", "4"}, {7, "Hearts", "7"},
		{3, "Clubs", "3"},
	}

	total := hand.CheckTotal()
	if total != 26 {
		t.Errorf("The total should be 26. Instead, we got: %d", total)
	}
	return
}
