package main

import (
	"github.com/HsiaoCz/code-beast/templfiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userCtrl := controllers.NewUserControl()

	app.Get("/user/show", userCtrl.HandleUserShow)

	app.Listen(":3002")
}
