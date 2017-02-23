package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

// Hand - the array of cards a player is holding. Max hand size will be eleven card.
type Hand []Card

// Protomeld - not quite a full meld, but container for candidate melds.
type Protomeld []Card

// Meld - Rummy standard melds. Must be at least three sequential cards of the
// same suit (ex. 2H-3H-4H) or at least three cards of same rank.
type Meld [][]Card

// Melds - array of possible configuration of melds in a player's hand.
type Melds [][][]Card

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

func (h *Hand) CheckUnmeldedCards() {
	return
}

// CheckTotal - checks the total number of points in a player's hand. It must be
// less than 10 to knock.
func (h *Hand) CheckTotal() (total int) {
	totals := []int{}
	melds := h.CheckMelds()

SEARCH:
	for _, card := range *h {
		total = 0
		for _, i := range melds {
			for _, j := range i {
				for _, k := range j {
					if card == k {
						continue SEARCH
					}
				}
			}
			total += card.value
		}
		totals = append(totals, total)
	}

	for _, min := range totals {
		if total > min {
			total = min
		} else if total == 0 {
			total = min
		}
	}
	return
}

// CheckMelds - gets all the possible melds that can be created with a hand.
func (h *Hand) CheckMelds() (melds Melds) {
	seq := h.CheckMeldSeqFirst()
	mult := h.CheckMeldMultFirst()

	if reflect.DeepEqual(seq, mult) {
		melds = append(melds, seq)
		return
	}

	melds = append(melds, seq, mult)
	return
}

// PrettyPrintMelds - pretty prints the melds and makes it readable.
func (m Melds) PrettyPrintMelds() (melds string) {
	if len(m) == 0 {
		return "No melds in hand."
	}

	for index, i := range m {
		melds += "Meld " + strconv.Itoa(index+1) + ": "
		for _, j := range i {
			for _, k := range j {
				melds += k.rank + k.suit[:1] + " "
			}
		}
		melds += "\n"
	}
	return
}

// SubsetOfMeld - checks if a meld being made is a subset of a previous meld
// made. Ex. 2C 3C 4C is a subset of 2C 3C 4C 5C 6C.
func (m Protomeld) SubsetOfMeld(melds Meld) bool {
	for _, ms := range melds {
		for i := 3; i <= len(ms); i++ {
			ms = ms[:i]
			if reflect.DeepEqual(m, ms) {
				return true
			}
		}
	}
	return false
}

// CheckMeldSeqFirst - checks the melds that can be made in the player's hand.
// This configuration checks for sequential melds first and does not store
// repeat cards.
func (h *Hand) CheckMeldSeqFirst() Meld {
	melds := Meld{}

	sort.Sort(ByValue(*h))
	for i := 0; i < len(*h); i++ {
		// Create a proto-meld. Since our hand is sorted by value,
		// we can do a linear search and check by value to complete a meld.
		meld := Protomeld{(*h)[i]}
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
			if !meld.SubsetOfMeld(melds) {
				melds = append(melds, meld)
			}
		}
	}

	for i := 0; i < len(*h); i++ {
		meld := Protomeld{(*h)[i]}
	SEARCH:
		for j := i + 1; j < len(*h); j++ {
			// Search for the cards of the same value.
			if (*h)[i].value == (*h)[j].value {
				for _, m := range melds {
					for _, n := range m {
						if n == (*h)[j] {
							continue SEARCH
						}
					}
				}
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

// CheckMeldMultFirst -checks the melds that can be made in the player's hand.
// This configuration checks for card multiples melds first and does not store
// repeat cards.
func (h *Hand) CheckMeldMultFirst() Meld {
	melds := Meld{}

	sort.Sort(ByValue(*h))
	for i := 0; i < len(*h); i++ {
		meld := Protomeld{(*h)[i]}
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

	for i := 0; i < len(*h); i++ {
		// Create a proto-meld. Since our hand is sorted by value,
		// we can do a linear search and check by value to complete a meld.
		meld := Protomeld{(*h)[i]}
	SEARCH:
		for k, j := 1, i+1; j < len(*h); j++ {
			// Search for the adjacent cards of ascending order.
			if (*h)[i].value+k == (*h)[j].value {
				if (*h)[i].suit == (*h)[j].suit {
					for _, m := range melds {
						for _, n := range m {
							if n == (*h)[j] {
								continue SEARCH
							}
						}
					}
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
	return melds
}
