package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8081"
	TYPE = "tcp"
)

func main() {
	listen, err := net.Listen(TYPE, fmt.Sprintf("%s:%s", HOST, PORT))
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()
	buf := make([]byte, 1024)
	for {
		conn, err := listen.Accept()
		if err != nil {
			handleError(err)
		}

		n, err := conn.Read(buf)
		if err != nil {
			handleError(err)
		}
		fmt.Println(string(buf[:n]))
	}
}

func handleError(err error) {
	if err != nil {
		slog.Error("there is an error please check out the reason", "error", err)
		return
	}
}
