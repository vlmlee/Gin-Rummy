package main

import (
	"reflect"
	"testing"
)

func TestDeckShouldHaveUniqueCards(t *testing.T) {
	deck := InitializeDeck()
	result := CheckDups(deck)
	if result {
		t.Error("There is a duplicate card.")
	}
	return
}

func TestDeckShouldBeShuffled(t *testing.T) {
	deck := InitializeDeck()
	if reflect.DeepEqual(deck[:13], Deck{
		{1, "Clubs", "A"},
		{2, "Clubs", "2"},
		{3, "Clubs", "3"},
		{4, "Clubs", "4"},
		{5, "Clubs", "5"},
		{6, "Clubs", "6"},
		{7, "Clubs", "7"},
		{8, "Clubs", "8"},
		{9, "Clubs", "9"},
		{10, "Clubs", "10"},
		{11, "Clubs", "J"},
		{12, "Clubs", "Q"},
		{13, "Clubs", "K"},
	}) {
		t.Error("The deck is not shuffled correctly")
	}

	deck.Shuffle()
	if reflect.DeepEqual(deck, Deck{
		{1, "Clubs", "A"},
		{2, "Clubs", "2"},
		{3, "Clubs", "3"},
		{4, "Clubs", "4"},
		{5, "Clubs", "5"},
		{6, "Clubs", "6"},
		{7, "Clubs", "7"},
		{8, "Clubs", "8"},
		{9, "Clubs", "9"},
		{10, "Clubs", "10"},
		{11, "Clubs", "J"},
		{12, "Clubs", "Q"},
		{13, "Clubs", "K"},
	}) {
		t.Error("The deck is not shuffled correctly")
	}
	return
}

func TestDrawCards(t *testing.T) {
	deck := InitializeDeck()
	testHand := &Hand{}
	testHand.DrawCard(&deck)

	if len(*testHand) == 0 {
		t.Error("Failed to draw a card.")
	}

	testHand.DrawCard(&deck)

	if CheckDups(*testHand) {
		t.Error("There are duplicates in the hand.")
	}
	if len(deck) != 50 {
		t.Errorf("Deck does not have 50 cards left but %d", len(deck))
	}
	return
}

func TestDeckDealsTenCardsToPlayers(t *testing.T) {
	deck := InitializeDeck()
	p1 := &Player{"Michael", Hand{}}
	p2 := &Player{"AI", Hand{}}
	deck.Deal(p1, p2)

	if len(p1.Hand) != 10 {
		t.Error("Player one did not draw 10 cards!")
	}
	if len(p2.Hand) != 10 {
		t.Error("Player two did not draw 10 card!")
	}
	if CheckDups(p1.Hand) || CheckDups(p2.Hand) {
		t.Error("There are duplicate cards in the hands!")
	}
	if len(deck) != 32 {
		t.Errorf("Deck does not have 32 cards left but %d", len(deck))
	}
	return
}

// Checks for duplicate objects in a array that have Card types.
func CheckDups(arr []Card) bool {
	dups := map[Card]bool{}
	for _, card := range arr {
		if dups[card] == true {
			return true
		}
		dups[card] = true
	}
	return false
}
