package store

import (
	"context"

	"github.com/HsiaoCz/code-beast/motel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStorer interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, bson.M) ([]*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func MongoRoomStoreInit(client *mongo.Client, coll *mongo.Collection) *MongoRoomStore {
	return &MongoRoomStore{
		client: client,
		coll:   coll,
	}
}
