# Gin-Rummy

####The Gin Rummy card game implemented in Go.

This is my favorite two player card game of all time and Go has been a really great language to use so I decided to implement it in Go.

## How to Install

You can build the go package and run the compiled program with:

```
$ go build -o Rummy *.go
$ ./Rummy
```

## Game Format

The game format is the same as traditional Gin Rummy.

Cards are pretty printed in the format: `RANK-SUIT`. For example, `AH` means `Ace of Hearts`, `KS` is the `King of Spades`. Each turn, the player will be prompt with four options:

```
1. DRAW CARD FROM DECK
2. PICKUP CARD FROM STACK
3. CHECK MELDS IN HAND
4. CHECK POINTS IN HAND
```

You can then choose your actions by entering 1, 2, 3, or 4. The player will always goes first and there will be a message that logs what each player does during their turn.

To knock, check to see if the number of points in your hand is less than or equal to 10 and then knock by typing `Y` when you get prompted. You will not get prompted if the points in your hands are greater than 10.

### Some slight differences from regular Gin Rummy:

1. Jacks have values of 11, Queens 12, and Kings 13.

2. Knocking with zero points will net 20 extra points.

3. There is no big and small gin.

4. There is also no score you need to reach to complete the game. It will go on forever until you choose to not start a new game.

### The Dumb AI

The game has an AI that plays *greedily*. It's pretty simple and only seeks to knock first without memorizing any of the cards it has seen. So that means it looks for cards that it can meld and ignores cards that will make the hand worse. The AI will pick up cards from the stack that can pair with a card in its hand as long as it makes the hand *better*. Else, it will draw from the deck. If the card on top of the stack makes a meld, it will draw from the stack.

## License: MIT
