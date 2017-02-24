# Gin-Rummy

The Gin Rummy card game implemented in Go. This is my favorite two player card game of all time and Go has been a really great language to use so I decided to implement it in Go.

## How to Install

You can build the go package and then run the compiled program with:

```
go build -o Rummy *.go
./Rummy
```

## Game Format

Cards are pretty printed in the format: `RANK-SUIT`.

For example, `AH` means `Ace of Hearts`, `KS` is the `King of Spades`.

The player will always goes first, and there will be a message that logs what each player does during their turn.

### Some slight differences from regular Gin Rummy:

Jacks have values of 11, Queens 12, and Kings 13.

Knocking with 0 points will net 20 extra points. There is no big or small gin. There is also no score you need to reach to complete the game. It will go on forever until you choose to not start a new game.

### The Dumb AI

The game has an AI that plays *greedily*. It's pretty simple and only seeks to knock first without memorizing any of the cards it has seen. So that means it looks for cards that it can meld and ignores cards that will make the hand worse. The AI will pick up cards from the stack that can pair with a card in its hand as long as it makes the hand *better*. Else, it will draw from the deck. If the card on top of the stack makes a meld, it will draw from the stack.

## License: MIT
