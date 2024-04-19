package api

import (
	"context"
	"testing"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi  = "mongodb://localhost:27017"
	testDBName = "hotel-reservation-test"
)

type teststore struct {
	client *mongo.Client
	store  *store.Store
}

func setup(t *testing.T) *teststore {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		t.Fatal(err)
	}
	hotelStore := store.NewMongoHotelStore(client)

	return &teststore{
		client: client,
		store: &store.Store{
			User:    store.NewMongoUserStore(client),
			Hotel:   hotelStore,
			Room:    store.NewMongoRoomStore(client, hotelStore),
			Booking: store.NewMongoBookingStore(client),
		},
	}
}

func (ts *teststore) teardown(t *testing.T) {
	if err := ts.client.Database(testDBName).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}
