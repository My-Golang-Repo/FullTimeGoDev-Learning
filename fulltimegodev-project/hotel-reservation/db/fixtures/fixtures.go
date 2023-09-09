package fixtures

import (
	"context"
	"fmt"
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/fulltimegodev/hotel-reservation-nana/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func AddUser(store *db.Store, fn, ln string, admin bool) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParam{
		FirstName: fn,
		LastName:  ln,
		Email:     fmt.Sprintf("%s@%s.com", fn, ln),
		Password:  fmt.Sprintf("%s_%s", fn, ln),
	})

	if err != nil {
		log.Fatal(err)
	}

	user.IsAdmin = admin

	insertedUser, err := store.User.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	return insertedUser
}

func AddHotel(store *db.Store, name, address string, rating int, room []primitive.ObjectID) *types.Hotel {
	var roomIDS = room
	if room == nil {
		roomIDS = []primitive.ObjectID{}
	}
	hotel := &types.Hotel{
		Name:    name,
		Address: address,
		Rooms:   roomIDS,
		Rating:  rating,
	}

	insertedHotel, err := store.Hotel.Insert(context.Background(), hotel)
	if err != nil {
		log.Fatal(err)
	}

	return insertedHotel
}

func AddRoom(store *db.Store, size string, ss bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		Seaside: ss,
		Price:   price,
		HotelID: hotelID,
	}

	insertedRoom, err := store.Room.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}

	return insertedRoom
}

func AddBooking(store *db.Store, userID, roomID primitive.ObjectID, from, till time.Time, numOfPerson int) *types.Booking {
	booking := &types.Booking{
		UserID:       userID,
		RoomID:       roomID,
		FromDate:     from,
		TillDate:     till,
		NumOfPersons: numOfPerson,
	}

	insertedBooking, err := store.Booking.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return insertedBooking
}
