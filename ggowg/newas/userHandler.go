package main

import "net/http"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) HandlerCreateUser(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, "some thing you need to check out")
}
