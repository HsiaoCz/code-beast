package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	lock      sync.RWMutex
	consumech chan RPC
	addr      NetAddr
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumech: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumech
}

func (t *LocalTransport) Connect(tr *LocalTransport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr
	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s could not send message to %s", t.addr, to)
	}

	peer.consumech <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
