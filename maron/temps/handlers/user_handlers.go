package handlers

import (
	"github.com/HsiaoCz/code-beast/crazy/templs/models"
	"github.com/HsiaoCz/code-beast/crazy/templs/views/userv"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func (h *UserHandler) HandleUserShow(c *fiber.Ctx) error {
	user := models.User{
		Username: "bob",
		Email:    "a@gg.com",
	}
	return Render(c, userv.UserShow(user))
}
