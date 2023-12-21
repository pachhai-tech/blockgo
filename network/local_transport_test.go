package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	// Perform type assertion to access LocalTransport's fields
	traLocal, ok1 := tra.(*LocalTransport)
	assert.True(t, ok1, "tra is not of type *LocalTransport")

	trbLocal, ok2 := trb.(*LocalTransport)
	assert.True(t, ok2, "trb is not of type *LocalTransport")

	assert.Equal(t, traLocal.peers[trb.Addr()], trb)
	assert.Equal(t, trbLocal.peers[tra.Addr()], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	tra.SendMessage(trb.Addr(), msg)

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.From, tra.Addr())
	assert.Equal(t, rpc.Payload, msg)
}
