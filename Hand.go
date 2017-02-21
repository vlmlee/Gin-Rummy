package main

import (
	"fmt"
	"sort"
)

// Hand is the array of cards a player is holding. Max hand size will be ten card.
type Hand []Card

// ByValue container for cards sorted by value.
type ByValue Hand

// BySuit container for cards sorted by suit.
type BySuit Hand

// PickUpAble describes a type that can add a card to player's hand. These
// types are Deck and Stack.
type PickUpAble interface {
	Len() int
	Pop() Card
	DrawCard() Card
}

// ByValue implemented by Sort
func (hand ByValue) Len() int {
	return len(hand)
}
func (hand ByValue) Less(i, j int) bool {
	return hand[i].value < hand[j].value
}
func (hand ByValue) Swap(i, j int) {
	hand[i], hand[j] = hand[j], hand[i]
}

// BySuit implemented by Sort
func (hand BySuit) Len() int {
	return len(hand)
}
func (hand BySuit) Less(i, j int) bool {
	return hand[i].suit < hand[j].suit
}
func (hand BySuit) Swap(i, j int) {
	hand[i], hand[j] = hand[j], hand[i]
}

func (s *Stack) Len() int {
	return len(*s)
}

func (d *Deck) Len() int {
	return len(*d)
}

func (s *Stack) Pop() Card {
	return (*s)[:len(*s)-1], (*s)[len(*s)]
}

// PrettyPrint a player's hand. This is for the view.
func (hand Hand) PrettyPrint() (result string) {
	// First sort Cards then pretty print
	sort.Sort(ByValue(hand))
	for i, card := range hand {
		result += card.symbol + card.suit[:1]
		if i != len(hand)-1 {
			result += " "
		}
	}
	return
}

// String() allows us to pretty print everytime we pass it to fmt.Print.
func (hand *Hand) String() string {
	return hand.PrettyPrint()
}

// CheckTotal checks the total number of points in a player's hand. It must be // less than 10 to knock.
func (hand *Hand) CheckTotal() (total int) {
	hand.CheckMeld()
	// for i, card := range *hand {

	// }
	return
}

// CheckMeld checks the melds that can be made in the player's hand. There may // be more than one meld configuration for various hands.
func (hand *Hand) CheckMeld() Hand {
	sort.Sort(ByValue(*hand))
	sort.Sort(BySuit(*hand))
	return *hand
}

// DrawCard by popping a card from a pickupable and appending it to a player's hand.
func (hand *Hand) DrawCard(pickupable PickUpAble) (d *PickUpAble, err error) {
	p := pickupable
	if len(*hand) >= 11 {
		err = fmt.Errorf("cannot have a hand size more than 11")
		return
	}
	card := p.Len()
	p = p[:len(p)-1]
	*hand = append(*hand, card)
	return
}

// DiscardCard places a card on top of the stack.
func (hand *Hand) DiscardCard(stack *Stack) (err error) {
	return
}
