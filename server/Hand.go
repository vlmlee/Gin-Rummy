package main

import (
	"fmt"
	"sort"
)

// Hand - the array of cards a player is holding. Max hand size will be eleven card.
type Hand []Card

// ByValue - container for cards sorted by value.
type ByValue Hand

// BySuit - container for cards sorted by suit.
type BySuit Hand

// PickUpAble - describes a type that can add a card to player's hand. These
// types are Deck and Stack.
type PickUpAble interface {
	DrawCard() Card
	IsEmpty() bool
}

// HasCard - describes a type that can look and see if it has a certain card.
type HasCard interface {
	ContainsCard() bool
}

// ByValue implements Sort method
func (hand ByValue) Len() int {
	return len(hand)
}
func (hand ByValue) Less(i, j int) bool {
	return hand[i].value < hand[j].value
}
func (hand ByValue) Swap(i, j int) {
	hand[i], hand[j] = hand[j], hand[i]
}

// BySuit implements Sort method
func (hand BySuit) Len() int {
	return len(hand)
}
func (hand BySuit) Less(i, j int) bool {
	return hand[i].suit < hand[j].suit
}
func (hand BySuit) Swap(i, j int) {
	hand[i], hand[j] = hand[j], hand[i]
}

// PrettyPrintHand - pretty prints a player's hand. This is for the view.
func (h Hand) PrettyPrintHand() (result string) {
	// First sort Cards then pretty print
	sort.Sort(ByValue(h))
	for i, card := range h {
		result += card.rank + card.suit[:1]
		if i != len(h)-1 {
			result += " "
		}
	}
	return
}

// String - allows us to pretty print everytime we pass it to fmt.Print.
func (h *Hand) String() string {
	return h.PrettyPrintHand()
}

// DrawCard by popping a card from a pickupable and appending it to a player's
// hand.
func (h *Hand) DrawCard(p PickUpAble) (cardDrawn Card, err error) {
	if len(*h) >= 11 {
		err = fmt.Errorf("cannot have a hand size more than 11")
		return
	}

	if p.IsEmpty() {
		err = fmt.Errorf("%T has no more cards", p)
		return
	}

	cardDrawn = p.DrawCard()
	*h = append(*h, cardDrawn)
	return
}

// DiscardCard - places a card on top of the stack.
func (h *Hand) DiscardCard(card Card, stack *Stack) (Card, error) {
	for i, cardInHand := range *h {
		if card == cardInHand {
			*h = append((*h)[:i], (*h)[i+1:]...)
			*stack = append((*stack), card)
			return card, nil
		}
	}
	return card, fmt.Errorf("could not find this card in hand")
}

// ContainsCard - checks and sees if the card is in the player's hand.
func (h *Hand) ContainsCard(card Card) bool {
	for _, cardInHand := range *h {
		if cardInHand == card {
			return true
		}
	}
	return false
}

// CheckTotal - checks the total number of points in a player's hand. It must be
// less than 10 to knock.
func (h *Hand) CheckTotal() (total int) {
	cardsInDeadwood := h.CheckDeadwood()
	for _, card := range cardsInDeadwood {
		total += card.value
	}
	return
}
