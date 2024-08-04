package main

import (
	"context"
	"fmt"
	"log"

	"github.com/qiniu/qmgo"
)

func main() {
	client, err := qmgo.NewClient(context.Background(), &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(qmgo.NewObjectID().Timestamp().UnixMilli()); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", client)
}
