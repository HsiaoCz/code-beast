package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

const CtxUserKey = "userID"

func JWTAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.GetReqHeaders()["Autherization"]
		if len(authHeader) == 0 {
			return errors.New("user unlogin")
		}
		if !(len(authHeader) == 2 && authHeader[0] == "Bearer") {
			return errors.New("invalid autherization")
		}
		mc, err := ParseToken(authHeader[1])
		if err != nil {
			return errors.New("invalid autherization")
		}
		c.Context().SetUserValue(CtxUserKey, mc.UserID)
		return c.Next()
	}
}
