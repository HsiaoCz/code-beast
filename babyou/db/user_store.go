package db

import (
	"context"

	"github.com/HsiaoCz/code-beast/babyou/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbname string, coll string) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(coll),
	}
}

func (mu *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	_, err := mu.coll.Find(ctx, bson.M{"email": user.Email})
	if err == nil {
		return nil, err
	}
	result, err := mu.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}
