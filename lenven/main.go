package main

import (
	"context"
	"log/slog"

	"github.com/HsiaoCz/code-beast/lenven/handlers"
	"github.com/HsiaoCz/code-beast/lenven/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clien, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(store.DBURL))
	if err != nil {
		slog.Error("connect to the mongo db error", "err", err)
		return
	}

	var (
		userStore = store.NewDefaultUserStore(clien)

		store = &store.Store{UserStore: userStore}

		userHandler = handlers.NewUserHandler(store)
	)

	app := fiber.New()
	v1 := app.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.Post("/create", userHandler.CreateUser)
		}
	}
	app.Listen(":9001")
}
