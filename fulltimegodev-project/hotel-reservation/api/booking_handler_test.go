package api

import (
	"fmt"
	"github.com/fulltimegodev/hotel-reservation-nana/db/fixtures"
	"testing"
	"time"
)

func TestAdminGetBookings(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	user := fixtures.AddUser(tdb.store, "Apollo", "Norm", false)
	hotel := fixtures.AddHotel(tdb.store, "UNCF Hotel", "JAPAN", 5, nil)
	room := fixtures.AddRoom(tdb.store, "King Size", true, 99.99, hotel.ID)
	booking := fixtures.AddBooking(tdb.store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2), 3)
	fmt.Println(booking)
	//app := fiber.New()
	//bookingHandler := NewBookingHandler(tdb.store)
	//app.Get("/booking", bookingHandler.HandleGetBookings)

}
