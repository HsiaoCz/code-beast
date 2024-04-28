package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/thstar/views/home"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Home())
}
