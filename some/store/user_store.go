package store

import (
	"context"
	"errors"

	"github.com/HsiaoCz/code-beast/some/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(DBUSERCOLL),
	}
}

func (m *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	// check the user exists
	_, err := m.coll.Find(ctx, bson.M{
		"email": user.Email,
	})
	if err == nil {
		return nil, errors.New("the user is exists")
	}
	res, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID).String()
	return user, nil
}
