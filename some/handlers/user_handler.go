package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/some/types"
	"github.com/HsiaoCz/code-beast/some/utils"
	"github.com/HsiaoCz/code-beast/some/views"
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

func (u *UserHandler) HandleShowUser(w http.ResponseWriter, r *http.Request) error {
	user := types.User{
		ID:       "1222333",
		Username: "bob",
		Email:    "gg@gg.com",
	}
	// maybe have little over staff
	return utils.Render(views.Show(user), w, r)
}
