package main

// Hand is the array of cards a player is holding. Max hand size is ten card.
type Hand struct {
	Cards       []Card
	MaxHandSize int
	HandSize    int
}

// AddCardToHand will pop a card from the deck and add it to the player's hand.
func (h *Hand) AddCardToHand(card []Card) Hand {
	return
}

// DiscardCard will pop a card from the player's hand and add it to the stack.
func (h *Hand) DiscardCard() Hand {

}

// PrettyPrint a player's hand. This is for the view.
func PrettyPrint(h Hand) string {
	var result string
	for i, card := range h.Cards {
		result += card.symbol + card.suit[:1]
		if i != len(h.Cards)-1 {
			result += " "
		}
	}
	return result
}

// String() allows us to pretty print everytime we pass it to fmt.Print.
func (h Hand) String() string {
	return PrettyPrint(h)
}

// CheckTotal checks the total number of points in a player's hand. It must be // less than 10 to knock.
func CheckTotal() {

}

// CheckMeld checks the melds that can be made in the player's hand. There may // be more than one meld configuration for various hands.
func CheckMeld() {

}
