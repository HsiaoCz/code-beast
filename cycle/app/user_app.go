package app

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/cycle/contact"
)

type UserHandler struct {
	con *contact.Dependenec
}

func NewUserHandler(con *contact.Dependenec) *UserHandler {
	return &UserHandler{
		con: con,
	}
}

func (u *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
