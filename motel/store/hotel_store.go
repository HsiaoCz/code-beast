package store

import (
	"context"

	"github.com/HsiaoCz/code-beast/motel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStorer interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(context.Context, bson.M, bson.M) error
	GetHotels(context.Context, bson.M) ([]*types.Hotel, error)
	GetHotelByID(context.Context, primitive.ObjectID) (*types.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func MongoHotelStoreInit(client *mongo.Client, coll *mongo.Collection) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   coll,
	}
}
