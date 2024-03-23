package main

import (
	"github.com/HsiaoCz/code-beast/slick"
	"github.com/HsiaoCz/code-beast/slick/app/profile"
	"github.com/HsiaoCz/code-beast/slick/app/view"
)

func main() {
	app := slick.New()
	app.Plugs(slick.WithAuth)
	app.Get("/hello", HandleHello)
	app.Get("/dashbord", HandleDashBord)
	app.Start(":9001")
}

func HandleHello(c *slick.Context) error {
	user := profile.User{
		FirstName: "anthony",
		LastName:  "gg",
		Email:     "anthdm@gg.com",
	}
	return c.Render(view.Index(user))
}

func HandleDashBord(c *slick.Context) error {
	return c.Render(view.Dashbord())
}
