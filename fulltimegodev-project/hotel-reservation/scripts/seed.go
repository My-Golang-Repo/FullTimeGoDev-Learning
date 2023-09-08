package main

import (
	"context"
	"fmt"
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/fulltimegodev/hotel-reservation-nana/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	userStore  db.UserStore
	ctx        = context.Background()
)

func main() {
	seedHotel("UNCF Hotel", "Japan", 5)
	seedHotel("Yamato Hotel", "Iscandar", 5)
	seedHotel("Andromeda", "Saturn", 5)
	seedUser(true, "admin@uncf.org", "admin", "uncf", "admin")
	seedUser(false, "achilles@uncf.org", "Achilles", "AAA-0004", "supersecurepassword")
}

func seedUser(isAdmin bool, email, fname, lname, password string) {
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

	_, err = userStore.InsertUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
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
}

func seedHotel(name string, address string, rating int) {
	hotel := types.Hotel{
		Name:    name,
		Address: address,
		Rooms:   []primitive.ObjectID{},
		Rating:  rating,
	}

	rooms := []types.Room{
		{
			Size:  "small",
			Price: 99.9,
		},
		{
			Size:  "normal",
			Price: 129.9,
		},
		{
			Size:  "kingsize",
			Price: 399.9,
		},
	}

	insertedHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedRoom)
	}
}
