package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(store.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := store.NewMongoHotelStore(client, store.DBNAME)
	roomStore := store.NewMongoRoomStore(client, store.DBNAME)
	hotel := types.Hotel{
		Name:      "Bellucia",
		Localtion: "France",
	}

	room := types.Room{
		Type:      types.SingleRoomType,
		BasePrice: 99.9,
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	room.HotelID = insertedHotel.ID
	insertRoom, err := roomStore.InsertRoom(ctx, &room)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertedHotel)
	fmt.Println(insertRoom)
}
