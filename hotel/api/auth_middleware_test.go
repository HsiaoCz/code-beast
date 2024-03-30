package api

import (
	"context"
	"testing"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type test_auth_store struct {
	store.UserStore
}

func setup_test_auth_store(t *testing.T) *test_auth_store {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(store.DBURI))
	if err != nil {
		t.Fatal(err)
	}
	return &test_auth_store{
		UserStore: store.NewMongoUserStore(client),
	}
}

func (ts *test_auth_store) teardown(t *testing.T) {
	if err := ts.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestToken(t *testing.T) {
	us_test_store := setup_test_auth_store(t)
	defer us_test_store.teardown(t)
	
	userHandler := NewUserHandler(us_test_store)
	app := fiber.New()
	app.Post("/api/v1/user", userHandler.HandlePostUser)
	app.Post("/api/v1/user/login", userHandler.HandleUserLogin)
	app.Listen(":9001")
}
