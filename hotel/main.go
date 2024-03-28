package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/HsiaoCz/code-beast/hotel/api"
	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"

// const dbname = "hotel-reservation"
// const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	},
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	// ctx := context.Background()

	// coll := client.Database(dbname).Collection(userColl)

	// user := types.User{
	// 	FirstName: "james",
	// 	LastName:  "At the water cooler",
	// }

	// res, err := coll.InsertOne(ctx, user)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("the response of mongo insert: %v\n", res)

	// var james types.User

	// if err := coll.FindOne(ctx, bson.M{}).Decode(&james); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("the james %+v\n", james)

	listenAddr := flag.String("listenAddr", ":9001", "set the listen address of the api server")
	flag.Parse()

	app := fiber.New(config)
	v1 := app.Group("/api/v1")
	{
		userHander := api.NewUserHandler(store.NewMongoUserStore(client))
		v1.Get("/user", userHander.HandleGetUsers)
		v1.Get("/user/:id", userHander.HandleGetUser)
		v1.Post("/user", userHander.HandlePostUser)
		v1.Delete("/user/:id", userHander.HandleDeleteUser)
		v1.Put("/user/:id", userHander.HandlePutUser)
	}
	app.Listen(*listenAddr)
}
