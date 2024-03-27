package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anthdm/hollywood/actor"
)

type Manager struct{}

func NewManager() actor.Producer {
	return func() actor.Receiver {
		return &Manager{}
	}
}

func (m *Manager) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Started:
	case actor.Stopped:
	}
}

func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	e.SpawnFunc(func(ctx *actor.Context) {
		switch msg := ctx.Message().(type) {
		case actor.Started:
			fmt.Println("started")
			_ = msg
		}
	}, "foo")
	e.Spawn(NewManager(), "manager")
	time.Sleep(time.Second * 2)
}
