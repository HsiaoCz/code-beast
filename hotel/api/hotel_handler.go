package api

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	hotelStore store.HotelStore
	roomStore  store.RoomStore
}

func NewHotelHandler(hotelStore store.HotelStore, roomStore store.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hotelStore,
		roomStore:  roomStore,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.hotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(hotels)
}
