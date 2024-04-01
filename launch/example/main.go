package main

import (
	"github.com/HsiaoCz/code-beast/launch"
	"github.com/HsiaoCz/code-beast/launch/example/handlers"
)

func main() {
	l := launch.New()
	l.Get("/user", handlers.CreateUser)
	l.Start(":9001")
}
