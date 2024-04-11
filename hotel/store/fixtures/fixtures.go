package fixtures

import (
	"context"
	"fmt"
	"log"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(store store.Store, email, fn, ln string, admin bool) *types.User {
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

func AddHotel(store store.Store, name string, loc string, rating int, rooms []primitive.ObjectID) *types.Hotel {

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
