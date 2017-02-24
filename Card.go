package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Card - type that has a suit, value, and rank (AJQK or numeric).
type Card struct {
	value int
	suit  string
	rank  string
}

// CreateDeckOfCards - does a linear map of suits and values and returns a deck
// of 52 cards.
func CreateDeckOfCards() (deck Deck) {
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	deck = make(Deck, 0)
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, GetCardWithRank(value, suit))
		}
	}
	return
}

// GetCardWithRank - appends the face value to the Card object.
func GetCardWithRank(value int, suit string) Card {
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

// PrettyPrintCard - pretty prints a card into "RANK-SUIT".
func (c Card) PrettyPrintCard() string {
	return c.rank + c.suit[:1]
}

func (c *Card) String() string {
	return c.PrettyPrintCard()
}

// GetCardFromPrettyPrint - transforms a card in the format "RANK-SUIT" into
// a card.
func GetCardFromPrettyPrint(p string) (card Card, err error) {
	var suit string
	var value int

	if p[:2] == "10" {
		suit, err = GetCardSuit(p[2:])
		if err != nil {
			err = fmt.Errorf("unidentified card suit")
		}

		value, err = GetCardValue(p[:2])
		if err != nil {
			err = fmt.Errorf("could not get the value of this card")
		}

		card = GetCardWithRank(value, suit)
		return
	}

	s := strings.Split(p, "")
	value, err = GetCardValue(s[0])
	suit, err = GetCardSuit(s[1])
	card = GetCardWithRank(value, suit)
	return
}

// GetCardSuit - gets the card's suit from the pretty printed string.
func GetCardSuit(s string) (suit string, err error) {
	switch s {
	case "C":
		suit = "Clubs"
	case "D":
		suit = "Diamonds"
	case "H":
		suit = "Hearts"
	case "S":
		suit = "Spades"
	default:
		return suit, fmt.Errorf("unidentified card suit")
	}
	return
}

// GetCardValue - gets the card's value from the pretty printed string.
func GetCardValue(s string) (value int, err error) {
	switch s {
	case "A":
		value = 1
	case "J":
		value = 11
	case "Q":
		value = 12
	case "K":
		value = 13
	case "10":
		value = 10
	default:
		value, err = strconv.Atoi(s)
		return value, err
	}
	return
}
