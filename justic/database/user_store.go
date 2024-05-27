package database

import (
	"context"

	"github.com/HsiaoCz/code-beast/justic/datastruction"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *datastruction.User) (*datastruction.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, coll *mongo.Collection) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoUserStore) CreateUser(ctx context.Context, user *datastruction.User) (*datastruction.User, error) {
	return nil, nil
}
