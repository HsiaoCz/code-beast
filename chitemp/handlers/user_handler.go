package handlers

import "net/http"

type UserHandler struct {
}

func NewUserHandler()*UserHandler{
	return &UserHandler{}
}

func (u *UserHandler) HandleShow(w http.ResponseWriter, r *http.Request) error {
	return nil
}
