package network

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestConnect(t *testing.T) {
	// assert.Equal(t, 1, 1)
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.peers[trb.addr], trb)
	assert.Equal(t, trb.peers[tra.addr], tra)
}


