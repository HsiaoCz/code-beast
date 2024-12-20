package main

import (
	"github.com/HsiaoCz/code-beast/maron/temps/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	userHandler := &handlers.UserHandler{}
	app.Get("/user", userHandler.HandleUserShow)
	app.Listen(":3001")
}
