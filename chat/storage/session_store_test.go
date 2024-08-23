package storage

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateSessions(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
}
