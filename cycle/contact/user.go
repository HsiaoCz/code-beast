package contact

import (
	"context"
	"errors"

	"github.com/HsiaoCz/code-beast/cycle/define"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserContact interface {
	CreateUser(context.Context, *define.User) (*define.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*define.User, error)
}

type MongoUserContact struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserContact(client *mongo.Client, coll *mongo.Collection) *MongoUserContact {
	return &MongoUserContact{
		client: client,
		coll:   coll,
	}
}

func (mc *MongoUserContact) CreateUser(ctx context.Context, user *define.User) (*define.User, error) {
	var result define.User
	filter := bson.D{{Key: "email", Value: user.Email}}
	if err := mc.coll.FindOne(ctx, filter).Decode(&result); err != mongo.ErrNoDocuments {
		return nil, errors.New("create user failed because this record created")
	}
	resp, err := mc.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = resp.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (mc *MongoUserContact) GetUserByID(ctx context.Context, uid primitive.ObjectID) (*define.User, error) {
	var user define.User
	filter := bson.D{{Key: "_id", Value: uid}}
	if err := mc.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
