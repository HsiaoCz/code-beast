package db

import (
	"context"

	"github.com/HsiaoCz/code-beast/commerce/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	GetUsers(context.Context) (*models.User, error)
}

type UserMongoStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewUserMongoStore(client *mongo.Client) *UserMongoStore {
	return &UserMongoStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(USERCOLL),
	}
}

func (u *UserMongoStore) GetUsers(ctx context.Context) (*models.User, error) {
	return nil, nil
}
