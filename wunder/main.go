package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	app := fiber.New()
	app.Get("/user/{id}", GetUserByID)
	app.Listen(":3001")
	fmt.Println("Jessica Did You sleep with Your Goddamn Teacher?")
}

func GetUserByID(c *fiber.Ctx) error {
	uid := c.Params("id")

	user := User{
		ID:       uid,
		Username: "Jessica",
		Email:    "Jessica@gmail.com",
	}
	return c.Status(http.StatusOK).JSON(&user)
}
