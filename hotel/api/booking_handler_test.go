package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/HsiaoCz/code-beast/hotel/store/fixtures"
)

func TestAdminGetBookings(t *testing.T) {
	store := setup(t)
	defer store.teardown(t)

	user := fixtures.AddUser(store.store, "james@james.com", "james", "foo", false)
	hotel := fixtures.AddHotel(store.store, "bar hotel", "a", 4, nil)
	room := fixtures.AddRoom(store.store, "small", true, 5.5, hotel.ID)

	booking := fixtures.AddBooking(store.store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))

	fmt.Println(booking)
}
