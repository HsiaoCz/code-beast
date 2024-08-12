package handlers

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestHandlerCreateUser(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
}
