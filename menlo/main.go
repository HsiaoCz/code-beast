package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/code-beast/menlo/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	var (
		addr   = os.Getenv("PORT")
		router = http.NewServeMux()
	)

	http.ListenAndServe(addr,router)
}
