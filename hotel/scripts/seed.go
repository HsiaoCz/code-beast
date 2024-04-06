package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/HsiaoCz/code-beast/hotel/api/middleware"
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
	IsAdmin    = true
	userStore  = store.NewMongoUserStore(client)
	wg         = &sync.WaitGroup{}
)

func seedUser(email string, fname string, lname string, password string) {
	user, err := types.NewUserFromPase(types.CreateUserParams{
		Email:     email,
		FirstName: fname,
		LastName:  lname,
		Password:  password,
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = IsAdmin
	_, err = userStore.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	u, err := userStore.GetUserByEmailAndPassword(context.Background(), types.AuthParams{Emial: user.Email, Password: user.EncryptedPasswrod})
	if err != nil {
		log.Fatal(err)
	}
	token, err := middleware.GenToken(u.ID, u.Email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", user.Email, token)
	wg.Done()
}

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
	wg.Done()
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
	wg.Add(4)

	go seedHotel("Bellucia", "France", 3)
	go seedHotel("The cozy hotel", "The Nederlands", 4)
	go seedHotel("Dont die in sleep", "london", 1)

	go seedUser("gg@ggc.com", "lisis", "assms", "sjsjkajs")
	wg.Wait()
}
