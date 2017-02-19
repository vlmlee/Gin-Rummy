package main

import (
	"strconv"
)

// Card is an type that has a suit, value, and symbol (AJQK or numeric).
type Card struct {
	value  int
	suit   string
	symbol string
}

// CreateDeckOfCards does a linear map of suits and values and returns a deck
// of 52 cards.
func CreateDeckOfCards() (deck Deck) {
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	deck = make([]Card, 0)
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, DetermineCardSymbol(value, suit))
		}
	}
	return
}

// DetermineCardSymbol appends the face value to the Card object.
func DetermineCardSymbol(value int, suit string) Card {
	switch value {
	case 1:
		return Card{value, suit, "A"}
	case 11:
		return Card{value, suit, "J"}
	case 12:
		return Card{value, suit, "Q"}
	case 13:
		return Card{value, suit, "K"}
	default:
		return Card{value, suit, strconv.Itoa(value)}
	}
}
