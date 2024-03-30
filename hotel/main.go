package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/HsiaoCz/code-beast/hotel/api"
	"github.com/HsiaoCz/code-beast/hotel/api/middleware"
	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(store.DBURI))
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

	// handlers init
	var (
		userStore  = store.NewMongoUserStore(client)
		userHander = api.NewUserHandler(userStore)

		hotelStore   = store.NewMongoHotelStore(client)
		roomStore    = store.NewMongoRoomStore(client, hotelStore)
		store        = &store.Store{User: userStore, Room: roomStore, Hotel: hotelStore}
		hotelHandler = api.NewHotelHandler(store)

		app = fiber.New(config)
		v1  = app.Group("/api/v1")
	)
	{
		// router
		// user
		v1.Get("/user", middleware.JWTAuthMiddleware(), userHander.HandleGetUsers)
		v1.Get("/user/:id", middleware.JWTAuthMiddleware(), userHander.HandleGetUser)
		v1.Post("/user", userHander.HandlePostUser)
		v1.Post("/user/login", userHander.HandleUserLogin)
		v1.Delete("/user/:id", middleware.JWTAuthMiddleware(), userHander.HandleDeleteUser)
		v1.Put("/user/:id", middleware.JWTAuthMiddleware(), userHander.HandlePutUser)

		// hotel router
		v1.Get("/hotel", middleware.JWTAuthMiddleware(), hotelHandler.HandleGetHotels)
		v1.Get("/hotel/:id/rooms", middleware.JWTAuthMiddleware(), hotelHandler.HandleGetRooms)
	}
	app.Listen(*listenAddr)
}
