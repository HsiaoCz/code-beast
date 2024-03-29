package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type teststore struct {
	store.UserStore
}

func setup(t *testing.T) *teststore {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(store.DBURI))
	if err != nil {
		t.Fatal(err)
	}
	return &teststore{
		UserStore: store.NewMongoUserStore(client),
	}
}

func (ts *teststore) teardown(t *testing.T) {
	if err := ts.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestPosUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userhandler := NewUserHandler(tdb.UserStore)
	app.Post("/user", userhandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "James",
		LastName:  "Foo",
		Password:  "jksmonfesniko122",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)
	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected lastname %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected Email %s but got %s", params.Email, user.Email)
	}

	fmt.Println(resp.Status)
}
