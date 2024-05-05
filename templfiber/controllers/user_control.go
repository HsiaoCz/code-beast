package controllers

import (
	"github.com/HsiaoCz/code-beast/templfiber/data"
	"github.com/HsiaoCz/code-beast/templfiber/template/show"
	"github.com/gofiber/fiber/v2"
)

type UserControl struct {
}

func NewUserControl() *UserControl {
	return &UserControl{}
}

func (u *UserControl) HandleUserShow(c *fiber.Ctx) error {
	user := data.User{
		Username: "gg",
		Email:    "gg@gg.com",
		ID:       "122233",
	}
	return Render(c, show.Show(user))
}
