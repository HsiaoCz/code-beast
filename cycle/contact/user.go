package contact

import (
	"context"

	"github.com/HsiaoCz/code-beast/cycle/define"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserContact interface {
	CreateUser(context.Context, *define.User) (*define.User, error)
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

func (mc *MongoUserContact) CreateUser(ctx context.Context, user *define.User) (*define.User, error){
	return nil,nil
}
