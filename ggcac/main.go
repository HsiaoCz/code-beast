package main

import (
	"log"
	"net"
	"time"

	"github.com/HsiaoCz/code-beast/ggcac/cache"
)

func main() {
	opts := ServerOpts{
		ListenAddr: ":49991",
		IsLeader:   true,
	}
	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", ":49991")
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("SET Foo Bar 2500"))
	}()
	server := NewServer(opts, cache.New())
	server.Start()
}
