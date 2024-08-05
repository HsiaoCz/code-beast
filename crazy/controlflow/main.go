package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{}

	msgch chan string
}

func NewServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string),
	}
}

func (s *Server) start() {
	fmt.Println("server starting")
	s.loop()
}

func (s *Server) loop() {
	for {
		select {
		case <-s.quitch:
			s.quit()
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
			fmt.Println("nothing hapend")
		}
	}
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received a message:", msg)
}

func (s *Server) quit() {
	fmt.Println("the server quit")
}

func main() {
	s := NewServer()
	go s.start()

	s.sendMessage("hey do this ....")

	time.Sleep(time.Second * 3)
}
