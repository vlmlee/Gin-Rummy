package main

// Hand is the array of cards a player is holding. Max hand size will be ten card.
type Hand []Card

// PickUpAble describes a type that can add a card to player's hand. These
// types are Deck and Stack.
type PickUpAble interface {
	DrawCard()
}

// PrettyPrint a player's hand. This is for the view.
func (hand Hand) PrettyPrint() (result string) {
	// First sort Cards then pretty print
	for i, card := range hand {
		result += card.symbol + card.suit[:1]
		if i != len(hand)-1 {
			result += " "
		}
	}
	return
}

// String() allows us to pretty print everytime we pass it to fmt.Print.
func (hand Hand) String() string {
	return hand.PrettyPrint()
}

// CheckTotal checks the total number of points in a player's hand. It must be // less than 10 to knock.
func CheckTotal() {

}

// CheckMeld checks the melds that can be made in the player's hand. There may // be more than one meld configuration for various hands.
func CheckMeld() {

}
