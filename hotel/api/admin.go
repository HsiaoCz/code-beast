package api

import (
	"errors"

	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	userInfo, ok := c.Context().UserValue(CtxUserInfo).(*types.UserInfo)
	if !ok {
		return errors.New("not authorized")
	}
	if !userInfo.IsAdmin {
		return errors.New("not authorized")
	}
	return c.Next()
}
