package p2p

import (
	"fmt"
	"io"
)

type Handler interface {
	HandleMessage(msg *Message) error
}

type DefaultHandler struct {
	Version string
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

func (h *DefaultHandler) HandleMessage(msg *Message) error {
	b, err := io.ReadAll(msg.Payload)
	if err != nil {
		return err
	}
	fmt.Printf("handling the msg from %s:%s", msg.From, string(b))
	return nil
}
