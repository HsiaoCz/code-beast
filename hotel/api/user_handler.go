package api

import (
	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	userStore store.UserStore
}

func NewUserHandler(userStore store.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errstr := params.Validate(); len(errstr) != 0 {
		return c.Status(http.StatusBadRequest).JSON(errstr)
	}
	user, err := types.NewUserFromPase(params)
	if err != nil {
		return err
	}

	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(insertedUser)
}
