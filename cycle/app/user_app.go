package app

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/code-beast/cycle/contact"
	"github.com/HsiaoCz/code-beast/cycle/define"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var userReq define.CreateRequestParam
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		return err
	}
	user, err := u.con.User.CreateUser(r.Context(), &define.User{Username: userReq.Username, Password: userReq.Password, Email: userReq.Email})
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, H{
		"message": "create user success!",
		"user":    user,
	})
}

func (u *UserHandler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	user, err := u.con.User.GetUserByID(r.Context(), uid)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, H{
		"message": "get user success!",
		"user":    user,
	})
}
