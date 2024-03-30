package store

import (
	"context"
	"fmt"

	"github.com/HsiaoCz/code-beast/hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const userColl = "users"

type Droper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Droper

	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error
	GetUserByEmailAndPassword(ctx context.Context, params types.AuthParams) (*types.UserLoginResponse, error)
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
	_, err := m.coll.Find(ctx, bson.M{"email": user.Email})
	if err == nil {
		return nil, err
	}
	res, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (m *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := m.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("not delete any user")
	}
	return nil
}

func (m *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {
	update := bson.D{
		{
			Key: "$set", Value: params.ToBSON(),
		},
	}
	_, err := m.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("----dropping user connection")
	return m.coll.Drop(ctx)
}

func (m *MongoUserStore) GetUserByEmailAndPassword(ctx context.Context, params types.AuthParams) (*types.UserLoginResponse, error) {
	enpwd, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return nil, err
	}
	var usersr types.UserLoginResponse
	if err := m.coll.FindOne(ctx, bson.M{"email": params.Emial, "EncryptedPassword": string(enpwd)}).Decode(&usersr); err != nil {
		return nil, err
	}
	return &usersr, nil
}
