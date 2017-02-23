package main

import "testing"

func TestCreatesCompleteDeck(t *testing.T) {
	testDeck := CreateDeckOfCards()
	deck := [52]Card{
		{1, "Clubs", "A"}, {2, "Clubs", "2"}, {3, "Clubs", "3"},
		{4, "Clubs", "4"}, {5, "Clubs", "5"}, {6, "Clubs", "6"},
		{7, "Clubs", "7"}, {8, "Clubs", "8"}, {9, "Clubs", "9"},
		{10, "Clubs", "10"}, {11, "Clubs", "J"}, {12, "Clubs", "Q"},
		{13, "Clubs", "K"}, {1, "Diamonds", "A"}, {2, "Diamonds", "2"},
		{3, "Diamonds", "3"}, {4, "Diamonds", "4"}, {5, "Diamonds", "5"},
		{6, "Diamonds", "6"}, {7, "Diamonds", "7"}, {8, "Diamonds", "8"},
		{9, "Diamonds", "9"}, {10, "Diamonds", "10"}, {11, "Diamonds", "J"},
		{12, "Diamonds", "Q"}, {13, "Diamonds", "K"}, {1, "Hearts", "A"},
		{2, "Hearts", "2"}, {3, "Hearts", "3"}, {4, "Hearts", "4"},
		{5, "Hearts", "5"}, {6, "Hearts", "6"}, {7, "Hearts", "7"},
		{8, "Hearts", "8"}, {9, "Hearts", "9"}, {10, "Hearts", "10"},
		{11, "Hearts", "J"}, {12, "Hearts", "Q"}, {13, "Hearts", "K"},
		{1, "Spades", "A"}, {2, "Spades", "2"}, {3, "Spades", "3"},
		{4, "Spades", "4"}, {5, "Spades", "5"}, {6, "Spades", "6"},
		{7, "Spades", "7"}, {8, "Spades", "8"}, {9, "Spades", "9"},
		{10, "Spades", "10"}, {11, "Spades", "J"}, {12, "Spades", "Q"},
		{13, "Spades", "K"},
	}

	for i := 0; i < len(testDeck); i++ {
		if testDeck[i] != deck[i] || len(testDeck) != 52 {
			t.Error("The deck is missing a card!")
		}
	}
	return
}

func TestPrettyPrintCard(t *testing.T) {
	card := Card{4, "Hearts", "4"}
	if card.PrettyPrintCard() != "4H" {
		t.Error("Card did not pretty print to 4H!")
	}
	return
}

func TestGetCardFromPrettyPrint(t *testing.T) {
	card, err := GetCardFromPrettyPrint("4H")
	if (card != Card{4, "Hearts", "4"}) {
		t.Error("Pretty printed 4H did not return a valid card!")
	}

	_, err = GetCardFromPrettyPrint("3E")
	if err == nil {
		t.Error("GetCardFromPrettyPrint did not catch suit error!")
	}

	card, err = GetCardFromPrettyPrint("KS")
	if (card != Card{13, "Spades", "K"}) {
		t.Error("Pretty printed KS did not return a valid card!")
	}
	return
}
