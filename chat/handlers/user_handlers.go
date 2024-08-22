package handlers

import "net/http"

type UserHandlers struct{}

func (u *UserHandlers) HandleUserLogin(w http.ResponseWriter, r *http.Request) error {
	return nil
}
