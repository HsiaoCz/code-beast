package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context)
}

type DefaultUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewDefaultUserStore(client *mongo.Client) *DefaultUserStore {
	return &DefaultUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(DBUSERCOLL),
	}
}

func (du *DefaultUserStore) CreateUser(ctx context.Context) {}
