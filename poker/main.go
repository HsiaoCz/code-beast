package main

import (
	"github.com/HsiaoCz/code-beast/poker/p2p"
)

func main() {
	// randn := rand.New(rand.NewSource(time.Now().UnixNano()))

	// for j := 0; j < 10; j++ {
	// 	d := deck.Shuffle(deck.New(), randn)
	// 	// card := deck.NewCard(deck.Spades, 1)
	// 	fmt.Println(d)
	// }
	cfg := p2p.ServerConfig{
		ListenAddr: ":3001",
	}
	server := p2p.NewServer(cfg)
	server.Start()
}
