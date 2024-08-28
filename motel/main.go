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

// const dbname = "hotel-reservation"
// const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if apiError, ok := err.(api.Error); ok {
			return c.Status(apiError.Code).JSON(apiError)
		}
		aerr := api.NewError(http.StatusInternalServerError, err.Error())
		return c.Status(aerr.Code).JSON(aerr)
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

		hotelStore     = store.NewMongoHotelStore(client)
		roomStore      = store.NewMongoRoomStore(client, hotelStore)
		bookingStore   = store.NewMongoBookingStore(client)
		store          = &store.Store{User: userStore, Room: roomStore, Hotel: hotelStore, Booking: bookingStore}
		hotelHandler   = api.NewHotelHandler(store)
		roomHandler    = api.NewRoomHandler(store)
		bookingHandler = api.NewBookingHandler(store)

		app   = fiber.New(config)
		v1    = app.Group("/api/v1")
		admin = app.Group("/api/v1/admin")
	)
	{
		// router
		// user
		v1.Get("/user", api.JWTAuthMiddleware(), userHander.HandleGetUsers)
		v1.Get("/user/:id", api.JWTAuthMiddleware(), userHander.HandleGetUser)
		v1.Post("/user", userHander.HandlePostUser)
		v1.Post("/user/login", userHander.HandleUserLogin)
		v1.Delete("/user/:id", api.JWTAuthMiddleware(), userHander.HandleDeleteUser)
		v1.Put("/user/:id", api.JWTAuthMiddleware(), userHander.HandlePutUser)

		// hotel router
		v1.Get("/hotel", api.JWTAuthMiddleware(), hotelHandler.HandleGetHotels)
		v1.Get("/hotel/:id", api.JWTAuthMiddleware(), hotelHandler.HandleGetHotelByID)
		v1.Get("/hotel/:id/rooms", api.JWTAuthMiddleware(), hotelHandler.HandleGetRooms)

		// room
		v1.Post("/room/:id/book", api.JWTAuthMiddleware(), roomHandler.HandleBookRoom)
		v1.Get("/room", api.JWTAuthMiddleware(), roomHandler.HandleGetRooms)

		// bookings handlers
		admin.Get("/booking", api.JWTAuthMiddleware(), bookingHandler.HandleGetBookings)
		v1.Get("/booking/:id", api.JWTAuthMiddleware(), bookingHandler.HandleGetBooking)
	}
	app.Listen(*listenAddr)
}
