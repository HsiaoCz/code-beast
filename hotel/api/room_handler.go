package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
	booking := &types.Booking{
		UserID:     userID,
		RoomID:     roomID,
		NumPersons: params.NumPersons,
	}
	if err := params.Validate(booking); err != nil {
		return err
	}
	where := bson.M{
		"fromDate": bson.M{
			"gte": booking.FromDate,
		},
		"tillDate": bson.M{
			"lte": booking.TillDate,
		},
	}
	bookings, err := r.store.Booking.GetBookings(c.Context(), where)
	if err != nil {
		return err
	}
	slog.Info("the bookings", "bookings", bookings)
	if len(bookings) > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"type":    "error",
			"message": fmt.Sprintf("the room %s already booked", c.Params("id")),
		})
	}
	insterd, err := r.store.Booking.InsertBooking(c.Context(), booking)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(insterd)
}

func (r *RoomHandler) HandleCancelRoom(c *fiber.Ctx) error {
	return nil
}
