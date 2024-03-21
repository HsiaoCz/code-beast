package main

import (
	"github.com/HsiaoCz/code-beast/sth/setempl/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	userHandler := handlers.UserHandler{}
	app.Get("/user", userHandler.HandleUserShow)
	app.Listen(":30001")
}
