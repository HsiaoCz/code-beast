package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/HsiaoCz/code-beast/hotel/store/fixtures"
	"github.com/HsiaoCz/code-beast/hotel/types"
	"github.com/gofiber/fiber/v2"
)

func TestAdminGetBookings(t *testing.T) {
	store := setup(t)
	defer store.teardown(t)

	user := fixtures.AddUser(store.store, "james@james.com", "james", "foo", false)
	hotel := fixtures.AddHotel(store.store, "bar hotel", "a", 4, nil)
	room := fixtures.AddRoom(store.store, "small", true, 5.5, hotel.ID)

	booking := fixtures.AddBooking(store.store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))

	admin := fixtures.AddUser(store.store, "admin@admin.com", "admin", "admin", true)

	_ = booking
	_ = admin

	app := fiber.New()
	bookingHandler := NewBookingHandler(store.store)
	app.Get("/", bookingHandler.HandleGetBookings)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", "111111")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 response got %d", resp.StatusCode)
	}
	var bookings []*types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}
	_ = bookings
}

func TestUserGetBooking(t *testing.T) {
	store := setup(t)
	defer store.teardown(t)

	user := fixtures.AddUser(store.store, "james@james.com", "james", "foo", false)
	hotel := fixtures.AddHotel(store.store, "bar hotel", "a", 4, nil)
	room := fixtures.AddRoom(store.store, "small", true, 5.5, hotel.ID)

	booking := fixtures.AddBooking(store.store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))

	admin := fixtures.AddUser(store.store, "admin@admin.com", "admin", "admin", true)

	_ = booking
	_ = admin

	app := fiber.New()
	_ = app

	// next time we fix it
}
