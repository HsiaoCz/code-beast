package network

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestConnect(t *testing.T) {
	// assert.Equal(t, 1, 1)
	tra := New("A")
	trb := New("B")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.peers[trb.addr], trb)
	assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestSendMessage(t *testing.T) {
	tra := New("A")
	trb := New("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("Hello World")
	assert.Nil(t, tra.SendMessage(tra.addr, msg))

	rpc := <-trb.Consume()

	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.addr)
}
