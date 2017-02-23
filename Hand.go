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
		result += card.symbol + card.suit[:1]
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
func (h *Hand) DrawCard(p PickUpAble) (err error) {
	if len(*h) >= 11 {
		return fmt.Errorf("cannot have a hand size more than 11")
	}
	*h = append(*h, p.DrawCard())
	return
}

// DiscardCard places a card on top of the stack.
func (h *Hand) DiscardCard(card Card, stack *Stack) (err error) {
	for i, v := range *h {
		if card == v {
			*h = append((*h)[:i], (*h)[i+1:]...)
			*stack = append((*stack), card)
			return
		}
	}
	return fmt.Errorf("could not find card in hand")
}

// CheckTotal - checks the total number of points in a player's hand. It must be
// less than 10 to knock.
func (h *Hand) CheckTotal() (total int) {
	h.CheckMeld()
	// for i, card := range *hand {

	// }
	return
}

// CheckMeld - checks the melds that can be made in the player's hand. There may
// be more than one meld configuration for various hands.
func (h *Hand) CheckMeld() [][]Card {
	melds := [][]Card{}

	sort.Sort(ByValue(*h))
	for i := 0; i < len(*h); i++ {
		// Create a proto-meld. Since our hand is sorted by value,
		// we can do a linear search and check by value to complete a meld.
		meld := []Card{(*h)[i]}
		for k, j := 1, i+1; j < len(*h); j++ {
			// Search for the adjacent cards of ascending order.
			if (*h)[i].value+k == (*h)[j].value {
				if (*h)[i].suit == (*h)[j].suit {
					meld = append(meld, (*h)[j])
					k++
				}
			}
		}
		// If the length of the protomeld is less than 3, then it's not a meld.
		if len(meld) >= 3 {
			melds = append(melds, meld)
		}
	}

	for i := 0; i < len(*h); i++ {
		meld := []Card{(*h)[i]}
		for j := i + 1; j < len(*h); j++ {
			// Search for the cards of the same value.
			if (*h)[i].value == (*h)[j].value {
				meld = append(meld, (*h)[j])
			}
		}
		// If the length of the protomeld is less than 3, then it's not a meld.
		if len(meld) >= 3 {
			melds = append(melds, meld)
		}
	}
	return melds
}
