package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/user/{id}", HandleGetUserByID)

	app.Listen(os.Getenv("PORT"))
}

func HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user := User{
		ID:       id,
		Username: "GG",
		Email:    "gg@gg.com",
	}

	return c.Status(http.StatusOK).JSON(&user)

}
