package handlers

import "net/http"

type UserHandlers struct{}

func UserHandlersInit() *UserHandlers {
	return &UserHandlers{}
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
