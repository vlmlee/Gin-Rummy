package main

import (
	"math/rand"
	"time"
)

// Deck is an array of Card objects.
type Deck []Card

// InitializeDeck - creates a deck of 52 cards and shuffle them.
func InitializeDeck() (deck Deck) {
	deck = CreateDeckOfCards()
	deck.Shuffle()
	return
}

// Shuffle - does a random swap of each element in the array.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range *d {
		r := rand.Intn(len(*d))
		(*d)[i], (*d)[r] = (*d)[r], (*d)[i]
	}
}

// Deal cards to player's hands
func (d *Deck) Deal(p1, p2 *Player) {
	count := 0
	for len(p1.Hand) < 10 || len(p2.Hand) < 10 {
		if count%2 == 0 {
			p1.Hand.DrawCard(d)
		} else {
			p2.Hand.DrawCard(d)
		}
		count++
	}
}

// DrawCard - picks up a card from the deck.
func (d *Deck) DrawCard() (card Card) {
	card = (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return
}

// IsEmpty - checks if deck is empty.
func (d *Deck) IsEmpty() bool {
	if len(*d) == 0 {
		return true
	}
	return false
}
