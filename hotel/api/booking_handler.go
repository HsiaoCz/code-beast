package api

import (
	"errors"
	"net/http"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingHandler struct {
	store *store.Store
}

func NewBookingHandler(store *store.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

// this needs to be admin authorized
func (b *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := b.store.Booking.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(bookings)
}

// this needs to be user authorized
func (b *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := b.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(booking)
}

func (b *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := b.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return err
	}
	userID, ok := c.Context().UserValue("userID").(primitive.ObjectID)
	if !ok {
		return errors.New("unauthorized")
	}
	if booking.UserID != userID {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"type":    "error",
			"message": "not authorized",
		})
	}
	// now the question is how do we cancel the booking
	if err := b.store.Booking.UpdateBooking(c.Context(), booking.ID.String(), bson.M{"canceled": true}); err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "cancel successed!",
	})
}
