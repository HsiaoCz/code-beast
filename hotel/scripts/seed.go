package main

import (
	"context"
	"log"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	roomStore  store.RoomStore
	hotelStore store.HotelStore
	ctx        = context.Background()
)

func seedHotel(name string, location string, rating int) {
	hotel := types.Hotel{
		Name:      name,
		Localtion: location,
		Rooms:     []primitive.ObjectID{},
		Rating:    rating,
	}

	rooms := []types.Room{
		{
			// Type:      types.SingleRoomType,
			Size:      "small",
			BasePrice: 99.9,
		},
		{
			// Type:      types.DeluxeRoomType,
			Size:      "nomal",
			BasePrice: 199.0,
		},
		{
			// Type:      types.SeaSideRoomType,
			Size:      "kingsize",
			BasePrice: 122.9,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(store.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(store.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore = store.NewMongoHotelStore(client)
	roomStore = store.NewMongoRoomStore(client, hotelStore)
}

func main() {
	seedHotel("Bellucia", "France", 3)
	seedHotel("The cozy hotel", "The Nederlands", 4)
	seedHotel("Dont die in sleep", "london", 1)
}
