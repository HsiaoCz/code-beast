package api

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
