package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/lenven/store"
	"github.com/HsiaoCz/code-beast/lenven/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	store *store.Store
}

func NewUserHandler(store *store.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errorstr := params.Validate(); len(errorstr) != 0 {
		return c.Status(http.StatusBadRequest).JSON(errorstr)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	userresp, err := u.store.UserStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(userresp)
}
