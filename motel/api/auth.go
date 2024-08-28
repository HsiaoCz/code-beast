package api

import (
	"errors"
	"strings"

	"github.com/HsiaoCz/code-beast/motel/types"
	"github.com/gofiber/fiber/v2"
)

const CtxUserInfo = "userInfo"

func JWTAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.GetReqHeaders()["Autherization"]
		if len(authHeader) == 0 {
			return errors.New("user unlogin")
		}
		authStr := authHeader[0]
		tokenStr := strings.Split(authStr, " ")
		if tokenStr[0] != "Bearer" {
			return errors.New("invalid token please use Bearer for prefix")
		}
		mc, err := ParseToken(tokenStr[1])
		if err != nil {
			return errors.New("invalid autherization")
		}
		userInfo := types.UserInfo{
			UserID:  mc.UserID,
			Email:   mc.Email,
			IsAdmin: mc.IsAdmin,
		}
		c.Context().SetUserValue(CtxUserInfo, &userInfo)
		return c.Next()
	}
}
