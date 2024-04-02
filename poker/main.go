package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/HsiaoCz/code-beast/poker/deck"
)

func main() {
	randn := rand.New(rand.NewSource(time.Now().UnixNano()))

	for j := 0; j < 10; j++ {
		d := deck.Shuffle(deck.New(), randn)
		// card := deck.NewCard(deck.Spades, 1)
		fmt.Println(d)
	}

}
