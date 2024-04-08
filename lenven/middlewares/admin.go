package middlewares

import (
	"errors"

	"github.com/HsiaoCz/code-beast/lenven/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return errors.New("not authorized")
	}
	if !user.IsAdmin {
		return errors.New("not authorized")
	}
	return c.Next()
}
