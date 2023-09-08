package api

import (
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/fulltimegodev/hotel-reservation-nana/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(bookingStore *db.Store) *BookingHandler {
	return &BookingHandler{
		store: bookingStore,
	}
}

// TODO: This should be admin authorized
func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.Booking.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(bookings)
}

// TODO: This needs to be user authorized
func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	booking, err := h.store.Booking.GetBookingByID(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}

	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return err
	}

	if booking.UserID != user.ID {
		return c.Status(http.StatusUnauthorized).JSON(genericResp{
			Type:    "error",
			Message: "Not Authorized, NOT AUTHORIZED",
		})
	}
	return c.JSON(booking)
}
