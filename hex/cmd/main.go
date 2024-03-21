package main

import (
	"fmt"
	"log"

	"github.com/HsiaoCz/code-beast/hex/internal/adapters/core"
)

func main() {
	adapter := core.NewAdapter()
	result, err := adapter.Addition(1, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", result)
}
