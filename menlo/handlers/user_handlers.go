package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/menlo/data"
)

type UserHandlers struct {
	UD data.UserDatar
}

func UserHandlersInit(UD data.UserDatar) *UserHandlers {
	return &UserHandlers{
		UD: UD,
	}
}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandlers) HandleUserLogin(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandlers) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandlers) HandleListUsers(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandlers) HandleDeleteUserByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandlers) HandleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
