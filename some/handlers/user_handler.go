package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/some/utils"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return utils.WriteJSON(w, http.StatusOK, utils.Map{
		"Message": "Hello",
	})
}
