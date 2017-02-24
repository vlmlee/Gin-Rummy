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

	err := hand.DrawCard(&deck)
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

func TestCheckMeld(t *testing.T) {
	hand1 := Hand{
		{13, "Clubs", "K"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {6, "Diamonds", "6"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {5, "Clubs", "5"}, {7, "Hearts", "7"},
		{11, "Clubs", "J"},
	}

	meld1 := hand1.CheckMelds()
	if reflect.DeepEqual(meld1, Meld{}) {
		t.Error("There were melds when there should've been none.")
	}

	hand2 := Hand{
		{13, "Clubs", "K"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {3, "Diamonds", "3"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {4, "Diamonds", "4"}, {7, "Hearts", "7"},
		{11, "Clubs", "J"},
	}

	meld2 := hand2.CheckMelds()
	if reflect.DeepEqual(meld2, Meld{
		{{2, "Diamonds", "2"}, {3, "Diamonds", "3"}, {4, "Diamonds", "4"}},
	}) {
		t.Error("The melds generated are incorrect. Expected 2D 3D 4D, got:", meld2)
	}

	hand3 := Hand{
		{2, "Clubs", "2"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {3, "Diamonds", "3"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {4, "Diamonds", "4"}, {7, "Hearts", "7"},
		{3, "Clubs", "3"},
	}

	meld3 := hand3.CheckMelds()
	if reflect.DeepEqual(meld3, Meld{
		{
			{1, "Clubs", "A"}, {2, "Clubs", "2"}, {3, "Clubs", "3"},
		},
		{
			{2, "Diamonds", "2"}, {3, "Diamonds", "3"}, {4, "Diamonds", "4"},
		},
		{
			{3, "Clubs", "3"}, {3, "Diamonds", "3"}, {3, "Spades", "3"},
		},
	}) {
		t.Error("The melds generated are incorrect.")
	}

	hand4 := Hand{
		{2, "Clubs", "2"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {3, "Diamonds", "3"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {4, "Diamonds", "4"}, {7, "Hearts", "7"},
		{5, "Diamonds", "5"},
	}

	meld4 := hand4.CheckMelds()
	if reflect.DeepEqual(meld4, Meld{
		{
			{2, "Diamonds", "2"}, {3, "Diamonds", "3"}, {4, "Diamonds", "4"},
			{5, "Diamonds", "5"},
		},
	}) {
		t.Error("The meld generated are incorrected. Expected 2D 3D 4D 5D.")
	}
	return
}

func TestPrettyPrintMeld(t *testing.T) {
	hand := Hand{
		{2, "Clubs", "2"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {3, "Diamonds", "3"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {4, "Diamonds", "4"}, {7, "Hearts", "7"},
		{3, "Clubs", "3"},
	}

	meld := hand.CheckMelds()
	if meld.PrettyPrintMelds() != "Meld 1: AC 2C 3C 2D 3D 4D \nMeld 2: 3C 3D 3S \n" {
		t.Error("Output for pretty print did not appear as expected.")
	}
	return
}

func TestPrettyPrintMeldNoMelds(t *testing.T) {
	hand := Hand{
		{13, "Clubs", "K"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {6, "Diamonds", "6"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {5, "Clubs", "5"}, {7, "Hearts", "7"},
		{11, "Clubs", "J"},
	}

	melds := hand.CheckMelds()
	if melds.PrettyPrintMelds() != "No melds in hand." {
		t.Errorf("Output should be 'No melds in hand'. Instead got: %s", melds.PrettyPrintMelds())
	}
}

func TestCheckUnmeldedCards(t *testing.T) {
	hand := Hand{
		{2, "Clubs", "2"}, {2, "Diamonds", "2"}, {1, "Clubs", "A"},
		{4, "Hearts", "4"}, {3, "Diamonds", "3"}, {12, "Spades", "Q"},
		{3, "Spades", "3"}, {4, "Diamonds", "4"}, {7, "Hearts", "7"},
		{3, "Clubs", "3"},
	}

	unmelds := hand.CheckUnmeldedCards()

	if reflect.DeepEqual(unmelds, []Unmelded{{
		{3, "Spades", "3"}, {4, "Hearts", "4"}, {7, "Hearts", "7"},
		{12, "Spades", "Q"}},
	}) {
		t.Error("Did not get the unmelded cards for the best meld.")
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
