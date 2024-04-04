package handlers

import (
	"github.com/HsiaoCz/code-beast/lenven/store"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	store *store.Store
}

func NewUserHandler(store *store.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	return nil
}
