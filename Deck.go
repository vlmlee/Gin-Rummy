package main

import (
	"fmt"
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
func (deck *Deck) Shuffle() (d Deck) {
	rand.Seed(time.Now().UTC().UnixNano())
	d = *deck
	for i := range d {
		r := rand.Intn(len(d))
		d[i], d[r] = d[r], d[i]
	}
	return
}

// Deal cards to player's hands
func (deck *Deck) Deal(p1, p2 *Player) {
	count := 0
	for len(p1.Hand) < 10 || len(p2.Hand) < 10 {
		if count%2 == 0 {
			*deck, _ = deck.DrawCard(&p1.Hand)
		} else {
			*deck, _ = deck.DrawCard(&p2.Hand)
		}
		count++
	}
}

// DrawCard by popping a card from the Deck and appending it to a player's hand.
func (deck *Deck) DrawCard(hand *Hand) (d Deck, err error) {
	d = *deck
	if len(*hand) >= 11 {
		err = fmt.Errorf("cannot have a hand size more than 11")
		return
	}
	card := d[len(d)-1]
	d = d[:len(d)-1]
	*hand = append(*hand, card)
	return
}
