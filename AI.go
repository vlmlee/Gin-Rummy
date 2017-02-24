package main

import "fmt"
import "math"

// AI can cheat and peek at all cards in the stack.

// The general strategy is to play greedily. The AI will ignore the player's
// actions and try to as quickly as possible to knock and end the game. This
// means it will simply look for nearest neighbors of cards in its hand that
// is currently not in a meld. If the card on the stack is not within 2 range,
// it will draw a card from the deck. If it has to decide between two cards
// with one pairing to discard, it will always discard the greater value card.

// CheckIfCardIsNearestNeighbor - will check if card on the top of the stack is
// a nearest neighbor.
func (h *Hand) CheckIfCardIsNearestNeighbor(c Card) bool {
	for _, card := range *h {
		if c.value == card.value+1 ||
			c.value == card.value-1 ||
			c.value == card.value {
			return true
		}
	}
	return false
}

// CardDistance - gives us the relative likelihood of the card being useful in
// our hand based on how far away it is from all the other cards in our hand.
func (u *Deadwood) CardDistance(c Card) int {
	distance := 0.0
	for _, card := range *u {
		distance += math.Abs(float64(card.value) - float64(c.value))
	}
	return int(distance)
}

// ChooseCardToDiscard will decide between cards to discard. The theory is that
// a good strategy is to get rid of the highest value card first in order to get
// closer to knocking. Since the game is a race to knock, we should try to
// the chances to do so. Saving slots for nearest neighbors in hopes of drawing
// the right card is not good unless we have an idea of what cards are going to
// appear. We could implement that in the future.
func (u *Deadwood) ChooseCardToDiscard() Card {
	cardToDiscard := Card{}
	max := 0
	for _, card := range *u {
		if card.value > max {
			max = card.value
			cardToDiscard = card
		}
	}
	return cardToDiscard
}

// AIActions - describes what the AI is going to do.
func AIActions(p *Player, deck *Deck, stack *Stack, knock *bool, draw *bool, gin *bool) {
	total := p.Hand.CheckTotal()

	if len(*deck) == 0 {
		*draw = true
		return
	}

	if p.Hand.CheckTotal() <= 10 {
		if total == 0 {
			*gin = true
		}
		*knock = true
		fmt.Printf("\nAI knocks!\n")
		return
	}

	// Sum total strategy.
	fauxHand := Hand{}
	cardFromTopOfStack, _ := GetCardFromPrettyPrint(stack.PeekAtStack())
	fauxHand = append(fauxHand, p.Hand...)
	fauxHand = append(fauxHand, cardFromTopOfStack)
	leftOverCards := fauxHand.CheckDeadwood()
	cardToDiscard := leftOverCards.ChooseCardToDiscard()
	umeldedCardsWithoutDraw := p.Hand.CheckDeadwood()
	deadWoodToDiscard := umeldedCardsWithoutDraw.ChooseCardToDiscard()

	// Nearest neighbor strategy
	nearestNeighbor := p.Hand.CheckIfCardIsNearestNeighbor(cardFromTopOfStack)

	if !nearestNeighbor {
		if cardToDiscard == cardFromTopOfStack {
			p.Hand.DrawCard(deck)
			fmt.Printf("\nAI drew from the deck!\n")
		} else if cardToDiscard.value < cardFromTopOfStack.value {
			p.Hand.DrawCard(deck)
			fmt.Printf("\nAI drew from the deck!\n")
		} else if leftOverCards.ContainsCard(cardFromTopOfStack) {
			p.Hand.DrawCard(deck)
			fmt.Printf("\nAI drew from the deck!\n")
		} else {
			p.Hand.DrawCard(stack)
			fmt.Printf("\nAI drew from the stack!\n")
		}
	} else {
		// If the card on the stack is worse than our worse card, then draw
		// from the deck.
		if umeldedCardsWithoutDraw.CardDistance(cardFromTopOfStack) > umeldedCardsWithoutDraw.CardDistance(deadWoodToDiscard) {
			p.Hand.DrawCard(deck)
			fmt.Printf("\nAI drew from the deck!\n")
			// Or if the card on the stack will become one of our deadwood
			// cards, we then draw from the deck.
		} else if leftOverCards.ContainsCard(cardFromTopOfStack) {
			p.Hand.DrawCard(deck)
			fmt.Printf("\nAI drew from the deck!\n")
		} else {
			p.Hand.DrawCard(stack)
			fmt.Printf("\nAI drew from the stack!\n")
		}
	}

	remainingCards := p.Hand.CheckDeadwood()
	discardedCard, _ := p.Hand.DiscardCard(remainingCards.ChooseCardToDiscard(), stack)
	fmt.Printf("%s discarded %s\n", p.name, discardedCard.PrettyPrintCard())
	return
}
