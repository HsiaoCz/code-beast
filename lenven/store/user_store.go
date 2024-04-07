package store

import (
	"context"
	"errors"

	"github.com/HsiaoCz/code-beast/lenven/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
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

func (du *DefaultUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	_, err := du.coll.Find(ctx, bson.M{"email": user.Email})
	if err == nil {
		return nil, errors.New("the user exists")
	}
	res, err := du.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID).String()
	return user, nil
}
