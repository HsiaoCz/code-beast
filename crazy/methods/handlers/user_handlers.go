package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/code-beast/crazy/methods/dao"
	"github.com/HsiaoCz/code-beast/crazy/methods/types"
)

type UserHandlers struct {
	user dao.UserCaser
}

func UserHandlersInit(user dao.UserCaser) *UserHandlers {
	return &UserHandlers{
		user: user,
	}
}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var create_user_params types.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&create_user_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	if err := create_user_params.Validate(); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	user, err := u.user.CreateUser(r.Context(), types.NewUserFromParams(create_user_params))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, user)
}

func (u *UserHandlers) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("uid")
	user, err := u.user.GetUserByID(r.Context(), id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return WriteJson(w, http.StatusOK, user)
}

func (u *UserHandlers) HandleDeleteUserByID(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("uid")
	if err := u.user.DeleteUserByID(r.Context(), id); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "delete user success",
	})
}
