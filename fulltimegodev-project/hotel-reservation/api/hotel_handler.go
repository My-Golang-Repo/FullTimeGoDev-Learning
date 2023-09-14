package api

import (
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type HotelQueryParams struct {
	db.Pagination
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var params HotelQueryParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}

	filter := db.Map{
		"rating": params.Rating,
	}

	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &params.Pagination)
	if err != nil {
		return ErrResourceNotFound("Hotels")
	}

	resp := ResourceResp{
		Results: len(hotels),
		Data:    hotels,
		Page:    int(params.Page),
	}

	return c.JSON(resp)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("Hotel")
	}
	return c.JSON(hotel)
}

func (h *HotelHandler) HandleGetHotelRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return ErrResourceNotFound("Rooms")
	}
	return c.JSON(rooms)
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
