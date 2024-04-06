package api

import (
	"errors"

	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
)

func GetAuthUser(c *fiber.Ctx) (*types.User, error) {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	return user, nil
}
