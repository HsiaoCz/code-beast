package api

import (
	"errors"
	"net/http"

	"github.com/HsiaoCz/code-beast/hotel/api/middleware"
	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"msg": "user not found"})
		}
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

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), userID); err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"delete": userID})
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		params types.UpdateUserParams
		userID = c.Params("id")
	)
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	if err := h.userStore.UpdateUser(c.Context(), filter, params); err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"updated": userID})
}

func (h *UserHandler) HandleUserLogin(c *fiber.Ctx) error {
	var authParams types.AuthParams
	if err := c.BodyParser(&authParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	userlv, err := h.userStore.GetUserByEmailAndPassword(c.Context(), authParams)
	if err != nil {
		return err
	}
	token, err := middleware.GenToken(userlv.ID, userlv.Email)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user":  userlv,
		"token": token,
	})
}
