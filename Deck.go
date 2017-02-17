package main

import (
	"math/rand"
	"time"
)

// Deck is an array of Card objects.
type Deck []Card

// InitializeDeck will create a deck of 52 cards and shuffle them.
func InitializeDeck() (deck Deck) {
	deck = CreateDeckOfCards()
	deck.Shuffle()
	return
}

// Shuffle does a random swap of each element in the array.
func (deck *Deck) Shuffle() Deck {
	rand.Seed(time.Now().UTC().UnixNano())
	d := *deck
	for i := range d {
		r := rand.Intn(len(d))
		d[i], d[r] = d[r], d[i]
	}
	return d
}
