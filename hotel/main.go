package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/HsiaoCz/code-beast/hotel/api"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	coll := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "james",
		LastName:  "At the water cooler",
	}

	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the response of mongo insert: %v\n", res)

	var james types.User

	coll.FindOne(ctx, bson.M{}).Decode(&james)

	listenAddr := flag.String("listenAddr", ":9001", "set the listen address of the api server")
	flag.Parse()

	app := fiber.New()
	v1 := app.Group("/api/v1")
	{
		v1.Get("/user", api.HandleGetUsers)
		v1.Get("/user/:id", api.HandleGetUser)
	}
	app.Listen(*listenAddr)
}
