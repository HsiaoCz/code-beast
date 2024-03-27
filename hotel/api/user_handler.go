package api

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "At the marshroom",
	}
	return c.Status(http.StatusOK).JSON(&u)
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "some user",
	})
}
