package store

import (
	"context"

	"github.com/HsiaoCz/code-beast/motel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Droper interface {
	Drop(context.Context) error
}

type UserStorer interface {
	Droper

	GetUserByID(context.Context, primitive.ObjectID) (*types.User, error)
	GetUser(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error
	GetUserByEmailAndPassword(ctx context.Context, params types.AuthParams) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func MongoUserStoreInit(client *mongo.Client, coll *mongo.Collection) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoUserStore) GetUserByID(ctx context.Context, id primitive.ObjectID) (*types.User, error) {
	var user types.User
	if err := m.coll.FindOne(ctx, bson.D{
		{Key: "_id", Value: id},
	}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	cur, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return []*types.User{}, err
	}
	if err := cur.All(ctx, &users); err != nil {
		return []*types.User{}, err
	}
	return users, nil
}
