package database

import (
	"context"
	"testing"
	"time"

	"github.com/HsiaoCz/code-beast/justic/datastruction"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoUrl = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.5"
	dbname   = "usertest"
	userColl = "users"
)

type Conn struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewConn() (*Conn, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	return &Conn{
		client: client,
		coll:   client.Database(dbname).Collection(userColl),
	}, nil
}

func TestCreateUser(t *testing.T) {
	conn, err := NewConn()
	if err != nil {
		t.Fatal(err)
	}
	users := []datastruction.User{
		{Username: "ben", Password: "123", Email: "ben@gmail.com"},
		{Username: "tom", Password: "234", Email: "tom@gmail.com"},
		{Username: "jessica", Password: "admin", Email: "jessica@gmail.com"},
	}

	var results []*datastruction.User
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for _, user := range users {
		result, err := CreateUser(conn, ctx, &user)
		if err != nil {
			t.Fatal(err)
		}
		results = append(results, result)
	}
	t.Logf("%v", results)

}

func CreateUser(conn *Conn, ctx context.Context, user *datastruction.User) (*datastruction.User, error) {
	// filter := bson.D{{Key: "email", Value: user.Email}}
	// result, err := conn.coll.Find(ctx, filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var users []datastruction.User
	// if err := result.All(ctx, &users); err != nil {
	// 	log.Fatal(err)
	// }
	res, err := conn.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID).String()
	return user, nil
}
