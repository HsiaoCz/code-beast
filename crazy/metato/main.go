package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Palyer struct {
	mu     sync.Mutex
	health int
}

func NewPlayer() *Palyer {
	return &Palyer{
		health: 100,
	}
}

func (p *Palyer) GetHealth() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.health
}

func (p *Palyer) TakeDamage(value int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.health -= value
}

func StartUILoop(p *Palyer) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Printf("player health:%d\r", p.GetHealth())
		<-ticker.C
	}
}

func StartGameLoop(p *Palyer) {
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		p.TakeDamage(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(40))
		if p.GetHealth() <= 0 {
			fmt.Println("Game Over")
			break
		}
		<-ticker.C
	}
}

func main(){
	player:=NewPlayer()
	go StartUILoop(player)
	StartGameLoop(player)
}