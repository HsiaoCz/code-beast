package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/justic/database"
)

type UserHandlers struct {
	store *database.Store
}

func NewUserHandlers(store *database.Store) *UserHandlers {
	return &UserHandlers{
		store: store,
	}
}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request)error{
	return WriteJSON(w,http.StatusOK,H{
		"Message":"all is well",
	})
}
