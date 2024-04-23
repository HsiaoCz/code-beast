package v1

import (
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

func (u *UserHandelrV1) HandleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := u.store.User.GetUsers(r.Context())
	if err != nil {
		slog.Error("get user from the db error", "err", err)
		return err
	}
	return WriteJSON(w, http.StatusOK, users)
}
