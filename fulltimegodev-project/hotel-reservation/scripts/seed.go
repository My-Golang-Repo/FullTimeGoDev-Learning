package main

import (
	"context"
	"fmt"
	"github.com/fulltimegodev/hotel-reservation-nana/api"
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/fulltimegodev/hotel-reservation-nana/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	client       *mongo.Client
	roomStore    db.RoomStore
	hotelStore   db.HotelStore
	userStore    db.UserStore
	bookingStore db.BookingStore
	ctx          = context.Background()
)

func main() {
	seedUser(true, "admin@uncf.org", "admin", "uncf", "admin")
	achilles := seedUser(false, "achilles@uncf.org", "Achilles", "AAA-0004", "supersecurepassword")
	hotel1 := seedHotel("Andromeda", "Saturn", 5)
	seedHotel("Yamato Hotel", "Iscandar", 5)
	seedHotel("UNCF Hotel", "Japan", 5)
	seedRoom("small", true, 99.99, hotel1.ID)
	seedRoom("medium", true, 199.99, hotel1.ID)
	room := seedRoom("kingsize", true, 399.99, hotel1.ID)
	seedBooking(achilles.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2), 2)
}

func seedUser(isAdmin bool, email, fname, lname, password string) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParam{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  password,
	})

	if err != nil {
		log.Fatal(err)
	}

	user.IsAdmin = isAdmin

	insertedUser, err := userStore.InsertUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s -> %s\n", user.Email, api.CreateTokenFromUser(user))

	return insertedUser
}

func seedHotel(name string, address string, rating int) *types.Hotel {
	hotel := types.Hotel{
		Name:    name,
		Address: address,
		Rooms:   []primitive.ObjectID{},
		Rating:  rating,
	}

	insertedHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	return insertedHotel
}

func seedRoom(size string, ss bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		Seaside: ss,
		Price:   price,
		HotelID: hotelID,
	}

	insertedRoom, err := roomStore.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}

	return insertedRoom
}

func seedBooking(userID, roomID primitive.ObjectID, from, till time.Time, numOfPersons int) {
	booking := &types.Booking{
		UserID:       userID,
		RoomID:       roomID,
		NumOfPersons: numOfPersons,
		FromDate:     from,
		TillDate:     till,
	}

	resp, err := bookingStore.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("booking id -> ", resp.ID)
}

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
	bookingStore = db.NewMongoBookingStore(client)
}
