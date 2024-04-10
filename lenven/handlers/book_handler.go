package handlers

import "github.com/gofiber/fiber/v2"

type BookHandler struct{}

func NewBookingHandler() *BookHandler {
	return &BookHandler{}
}

func (b *BookHandler) HandleCreateBooking(c *fiber.Ctx) error {
	return nil
}
