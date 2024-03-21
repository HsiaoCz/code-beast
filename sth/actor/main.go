package main

import "sync"

type Handler struct {
	lock  sync.Mutex
	state uint
}

func (h *Handler) handleMessage(msg uint) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.state = msg
}

func main() {
	h := &Handler{}

	for i := 0; i < 10; i++ {
		h.handleMessage(uint(i))
	}
}
