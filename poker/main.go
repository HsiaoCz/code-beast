package main

import (
	"fmt"

	"github.com/HsiaoCz/code-beast/poker/deck"
)

func main() {
	card := deck.NewCard(deck.Spades, 1)
	fmt.Println(card)
}
