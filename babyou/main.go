package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HsiaoCz/code-beast/babyou/db"
	"github.com/HsiaoCz/code-beast/babyou/handlers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoUrl = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.5"
	dbname   = "babyou"
	userColl = "users"
	port     = ":3001"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	var (
		mongoUserStore = db.NewMongoUserStore(client, dbname, userColl)
		store          = &db.Store{User: mongoUserStore}
		userHandler    = handlers.NewUserHandler(store)
	)
	fmt.Printf("%+v", client)
	app := fiber.New()

	app.Post("/user/create", userHandler.HandleCreateUser)

	app.Listen(port)
}
