package api

import (
	"net/http"
	"time"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct {
	store *store.Store
}

func NewRoomHandler(store *store.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

func (r *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params types.BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	roomID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}
	userID, ok := c.Context().Value("userID").(primitive.ObjectID)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"type":    "error",
			"message": "internal server error",
		})
	}
	booking := types.Booking{
		UserID:     userID,
		RoomID:     roomID,
		NumPersons: params.NumPersons,
	}
	formDate, err := time.Parse("2006-01-02", params.FromDate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"type":    "error",
			"message": "the formDate is unvalid",
		})
	}
	tillDate, err := time.Parse("2006-01-02", params.TillDate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"type":    "error",
			"message": "the tillDate is unvalid",
		})
	}
	booking.FromDate = formDate
	booking.TillDate = tillDate
	return c.Status(http.StatusOK).JSON(&booking)
}
