package fixtures

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(store *store.Store, email, fn, ln string, admin bool) *types.User {
	user, err := types.NewUserFromPase(types.CreateUserParams{
		Email:     email,
		FirstName: fn,
		LastName:  ln,
		Password:  fmt.Sprintf("%s_%s", fn, ln),
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = admin
	insertedUser, err := store.User.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func AddHotel(store *store.Store, name string, loc string, rating int, rooms []primitive.ObjectID) *types.Hotel {

	var roomIDS = rooms
	if rooms == nil {
		roomIDS = []primitive.ObjectID{}
	}

	hotel := types.Hotel{
		Name:      name,
		Localtion: loc,
		Rooms:     roomIDS,
		Rating:    rating,
	}
	insertedHotel, err := store.Hotel.InsertHotel(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func AddRoom(store *store.Store, size string, ss bool, price float64, hid primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		Seaside: ss,
		Price:   price,
		HotelID: hid,
	}
	insertRoom, err := store.Room.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertRoom
}

func AddBooking(store *store.Store, uid, rid primitive.ObjectID, from, till time.Time) *types.Booking {
	booking := &types.Booking{
		UserID:   uid,
		RoomID:   rid,
		FromDate: from,
		TillDate: till,
	}
	bking, err := store.Booking.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return bking
}
