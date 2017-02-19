package main

// Player is a type with a name and a Hand.
type Player struct {
	name string
	Hand
}

// DiscardCardFromHand will pop a card from the player's hand and add it to the // stack.
func (p *Player) DiscardCardFromHand(c Card, s *Stack) (hand Hand) {
	hand = p.Hand
	for index, card := range hand {
		if c == card {
			s.AddCardToStack(card)
			hand = append(hand[:index], hand[index+1:]...)
		}
	}
	return
}
