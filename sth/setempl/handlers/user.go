package handlers

import (
	"github.com/HsiaoCz/code-beast/sth/setempl/model"
	"github.com/HsiaoCz/code-beast/sth/setempl/view/userv"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c *fiber.Ctx) error {
	u := model.User{
		Email: "a@gg.com",
	}
	return Render(c, userv.Show(u))
}
