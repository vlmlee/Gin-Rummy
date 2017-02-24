package main

import (
	"reflect"
	"sort"
	"strconv"
)

// Protomeld - not quite a full meld, but container for candidate melds.
type Protomeld []Card

// Unmelded - cards that are not in a meld.
type Unmelded []Card

// Meld - Rummy standard melds. Must be at least three sequential cards of the
// same suit called a run (ex. 2H-3H-4H) or at least three cards of same rank called a set.
type Meld [][]Card

// Melds - array of possible configuration of melds in a player's hand.
type Melds [][][]Card

// ContainsCard - checks and sees if the card is in the unmelded set.
func (u Unmelded) ContainsCard(card Card) bool {
	for _, c := range u {
		if c == card {
			return true
		}
	}
	return false
}

// CheckUnmeldedCards - returns cards not in a meld configuration.
func (h *Hand) CheckUnmeldedCards() (unmelded Unmelded) {
	meldConfig := h.CheckMelds()
SEARCH_UNMELDED_CARDS:
	for _, cardInHand := range *h {
		for _, melds := range meldConfig {
			for _, meld := range melds {
				for _, cardInMeld := range meld {
					if cardInHand == cardInMeld {
						continue SEARCH_UNMELDED_CARDS
					}
				}
			}
			if !unmelded.ContainsCard(cardInHand) {
				unmelded = append(unmelded, cardInHand)
			}
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
	if reflect.DeepEqual(m, Melds{{}}) {
		return "No melds in hand."
	}

	for index, i := range m {
		melds += "Meld " + strconv.Itoa(index+1) + ": "
		for _, j := range i {
			for _, k := range j {
				melds += k.rank + k.suit[:1] + " "
			}
		}
		if index != len(m)-1 {
			melds += "\n"
		}
	}
	return
}

// SubsetOfMeld - checks if a meld being made is a subset of a previous meld
// made. Ex. 2C 3C 4C is a subset of 2C 3C 4C 5C 6C.
func (m Protomeld) SubsetOfMeld(melds Meld) bool {
	for _, ms := range melds {
		for i := 0; i+len(m) <= len(ms); i++ {
			subset := Protomeld{}
			subset = ms[i : i+len(m)]
			if reflect.DeepEqual(m, subset) {
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
			if !meld.SubsetOfMeld(melds) {
				melds = append(melds, meld)
			}
		}
	}
	return melds
}

// CheckMeldMultFirst - checks the melds that can be made in the player's hand.
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
			if !meld.SubsetOfMeld(melds) {
				melds = append(melds, meld)
			}
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
			if !meld.SubsetOfMeld(melds) {
				melds = append(melds, meld)
			}
		}
	}
	return melds
}
