package api

import (
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	hotelStore db.HotelStore
	roomStore  db.RoomStore
}

func NewMongoHotelStore(hotelStore db.HotelStore, roomStore db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hotelStore,
		roomStore:  roomStore,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.hotelStore.GetHotels(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}

func (h *HotelHandler) HandleInsertHotel(c *fiber.Ctx) error {
	//var params types.Hotel
	//
	//if err := c.BodyParser(&params); err != nil {
	//	return err
	//}
	//
	//insertedHotel, err := h.hotelStore.InsertHotel(c.Context(), &params)
	//if err != nil {
	//	return err
	//}
	return nil

}
