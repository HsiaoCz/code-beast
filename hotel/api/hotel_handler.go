package api

import (
	"net/http"

	"github.com/HsiaoCz/code-beast/hotel/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	// hotelStore store.HotelStore
	// roomStore  store.RoomStore
	store *store.Store
}

func NewHotelHandler(store *store.Store) *HotelHandler {
	return &HotelHandler{
		// hotelStore: hotelStore,
		// roomStore:  roomStore,
		store: store,
	}
}

// type HotelQueryParams struct {
// 	Rooms  bool
// 	Rating int
// }

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(rooms)
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	// var qparams HotelQueryParams

	// if err := c.QueryParser(&qparams); err != nil {
	// 	return err
	// }
	hotels, err := h.store.Hotel.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(hotels)
}

func (h *HotelHandler) HandleGetHotelByID(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), oid)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(hotel)
}

func (h *HotelHandler) HandleUpdateHotel(c *fiber.Ctx) error {
	return nil
}
