package handlers

import (
	"net/http"

	views "github.com/HsiaoCz/code-beast/chitemp/views/layout"
	"github.com/HsiaoCz/code-beast/chitemp/views/uservs"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) HandleShow(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Base(uservs.Show()))
}
