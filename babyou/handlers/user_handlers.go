package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/babyou/db"
	"github.com/HsiaoCz/code-beast/babyou/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	store *db.Store
}

func NewUserHandler(store *db.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var userc types.CreateUser
	if err := c.BodyParser(&userc); err != nil {
		return err
	}
	user := &types.User{
		Username: userc.Username,
		Password: userc.Password,
		Email:    userc.Email,
	}
	userrepy, err := u.store.User.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create user success",
		"user":    userrepy,
	})
}
