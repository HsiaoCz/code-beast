package main

import (
	"log"
	"net"
)

const (
	TYPE = "tcp"
	PORT = ":8081"
)

func main() {
	conn, err := net.Dial(TYPE, PORT)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn.Write([]byte{})
	}
}
