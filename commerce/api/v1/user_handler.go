package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/HsiaoCz/code-beast/commerce/db"
)

type UserHandelrV1 struct {
	store *db.Store
}

func NewUserHandlerV1(store *db.Store) *UserHandelrV1 {
	return &UserHandelrV1{
		store: store,
	}
}

func (u *UserHandelrV1) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.store.User.GetUsers(r.Context())
	if err != nil {
		slog.Error("get user from the db error", "err", err)
		return
	}
	json.NewEncoder(w).Encode(users)
}
