package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/babyou/db"
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
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "all is well",
	})
}
