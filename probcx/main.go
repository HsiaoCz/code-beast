package main

import (
	"time"

	"github.com/HsiaoCz/code-beast/probcx/network"
)

// Server
// Transport

func main() {

	trlocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trlocal.Connect(trRemote)
	trRemote.Connect(trlocal)

	go func() {
		trRemote.SendMessage(trlocal.Addr(), []byte("hello world"))
		time.Sleep(1 * time.Second)
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trlocal},
	}
	s := network.NewServer(opts)
	s.Start()
}
