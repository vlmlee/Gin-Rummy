package main

import (
	"reflect"
	"testing"
)

func TestDeckShouldHaveUniqueCards(t *testing.T) {
	deck := InitializeDeck()
	result := CheckDups(deck)
	if !result {
		t.Error("There is a duplicate card.")
	}
	return
}

func TestDeckShouldBeShuffled(t *testing.T) {
	deck := InitializeDeck()
	if reflect.DeepEqual(deck[:13], []Card{
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

	shuffledDeck := deck.Shuffle()
	if reflect.DeepEqual(shuffledDeck[:13], []Card{
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

// Checks for duplicate objects in an array.
func CheckDups(deck []Card) bool {
	dups := map[Card]bool{}
	for _, card := range deck {
		if dups[card] == true {
			return false
		}
		dups[card] = true
	}
	return true
}
