package api

import (
	"encoding/json"
	"fmt"
	"github.com/fulltimegodev/hotel-reservation-nana/db/fixtures"
	"github.com/fulltimegodev/hotel-reservation-nana/types"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAdminGetBookings(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	var (
		adminUser = fixtures.AddUser(tdb.store, "Admin", "Golang", true)
		user      = fixtures.AddUser(tdb.store, "Apollo", "Norm", false)
		hotel     = fixtures.AddHotel(tdb.store, "UNCF Hotel", "JAPAN", 5, nil)
		room      = fixtures.AddRoom(tdb.store, "King Size", true, 99.99, hotel.ID)
		booking   = fixtures.AddBooking(tdb.store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2), 3)

		app            = fiber.New(fiber.Config{ErrorHandler: ErrorHandler})
		admin          = app.Group("/", JWTAuthentication(tdb.store.User), AdminAuth)
		bookingHandler = NewBookingHandler(tdb.store)
	)

	admin.Get("/", bookingHandler.HandleGetBookings)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(adminUser))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expecting %d status code but got %d", http.StatusOK, resp.StatusCode)
	}

	var bookings []*types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}

	if len(bookings) != 1 {
		t.Fatalf("Expected booking to be 1 but got %d", len(bookings))
	}

	have := bookings[0]

	if have.ID != booking.ID {
		t.Fatalf("expected %s but got %s", booking.ID, have.ID)
	}

	if have.UserID != booking.UserID {
		t.Fatalf("expected %s but got %s", booking.UserID, have.UserID)
	}

	if have.RoomID != booking.RoomID {
		t.Fatalf("expected %s but got %s", booking.RoomID, have.RoomID)
	}

	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != ErrUnAuthorized().Code {
		t.Fatalf("expected status unauthorized %d responses but got %d", ErrUnAuthorized().Code, resp.StatusCode)
	}
}

func TestUserGetBooking(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	var (
		nonAuthUser = fixtures.AddUser(tdb.store, "Jimmy", "WaterCooler", false)
		user        = fixtures.AddUser(tdb.store, "Apollo", "Norm", false)
		hotel       = fixtures.AddHotel(tdb.store, "UNCF Hotel", "JAPAN", 5, nil)
		room        = fixtures.AddRoom(tdb.store, "King Size", true, 99.99, hotel.ID)
		booking     = fixtures.AddBooking(tdb.store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2), 3)

		app            = fiber.New()
		route          = app.Group("/", JWTAuthentication(tdb.store.User))
		bookingHandler = NewBookingHandler(tdb.store)
	)

	route.Get("/:id", bookingHandler.HandleGetBooking)

	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status code but got %d", http.StatusOK, resp.StatusCode)
	}

	var userBooking *types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&userBooking); err != nil {
		log.Fatal(err)
	}

	if userBooking.ID != booking.ID {
		log.Fatalf("expected booking to be %s but got %s", booking.ID, userBooking.ID)
	}

	if userBooking.UserID != booking.UserID {
		log.Fatalf("expected user ID to be %s but got %s", booking.UserID, userBooking.UserID)
	}

	// AKAN ERROR KARENA INI BOOKING KITA PAKAI USER ID BUKAN NON AUTH
	req = httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(nonAuthUser))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected status code to be non 200 but got %d", resp.StatusCode)
	}

}
