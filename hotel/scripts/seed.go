package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/HsiaoCz/code-beast/hotel/api"
	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/store/fixtures"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client       *mongo.Client
	roomStore    store.RoomStore
	hotelStore   store.HotelStore
	bookingStore store.BookingStore
	ctx          = context.Background()
	IsAdmin      = true
	userStore    = store.NewMongoUserStore(client)
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
	token, err := api.GenToken(u.ID, u.Email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", user.Email, token)
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
}

func seedRoom(size string, ss bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		Seaside: ss,
		Price:   price,
		HotelID: hotelID,
	}
	insertRoom, err := roomStore.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertRoom
}

func seedBooking(userID, roomID primitive.ObjectID, from, till time.Time) {
	booking := &types.Booking{
		UserID:   userID,
		RoomID:   roomID,
		FromDate: from,
		TillDate: till,
	}
	_, err := bookingStore.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
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
	userStore = store.NewMongoUserStore(client)
	bookingStore = store.NewMongoBookingStore(client)
}

func main() {
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
	userStore = store.NewMongoUserStore(client)
	bookingStore = store.NewMongoBookingStore(client)

	store := &store.Store{
		User:    userStore,
		Hotel:   hotelStore,
		Room:    roomStore,
		Booking: bookingStore,
	}
	user := fixtures.AddUser(store, "gg@gg.com", "final", "bob", false)
	fmt.Println("final---->", user.ID)

	admin := fixtures.AddUser(store, "admin@gg.com", "admin", "admin", true)
	fmt.Println("admin----->", admin.ID)

	hotel := fixtures.AddHotel(store, "some hotel", "bermude", 5, nil)
	fmt.Println(hotel)

	room := fixtures.AddRoom(store, "large", true, 88.44, hotel.ID)
	fmt.Println(room)

	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))
	fmt.Println(booking)

	seedHotel("Bellucia", "France", 3)
	seedHotel("The cozy hotel", "The Nederlands", 4)
	seedHotel("Dont die in sleep", "london", 1)
	seedRoom("small", true, 99.99, primitive.NewObjectID())
	seedBooking(primitive.NewObjectID(), primitive.NewObjectID(), time.Now(), <-time.After(time.Hour*24))
	seedUser("gg@ggc.com", "lisis", "assms", "sjsjkajs")
}
