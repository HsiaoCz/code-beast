package store

import (
	"context"

	"github.com/HsiaoCz/code-beast/hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) UserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(userColl),
	}
}

func (m *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	// validate the correctnes of the ID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user types.User
	if err := m.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
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

func (m *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}
