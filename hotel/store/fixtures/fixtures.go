package fixtures

import (
	"context"
	"fmt"
	"log"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
)

func AddUser(store store.Store, email, fn, ln string, admin bool) *types.User {
	user, err := types.NewUserFromPase(types.CreateUserParams{
		Email:     email,
		FirstName: fn,
		LastName:  ln,
		Password:  fmt.Sprintf("%s_%s", fn, ln),
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = admin
	insertedUser, err := store.User.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}
