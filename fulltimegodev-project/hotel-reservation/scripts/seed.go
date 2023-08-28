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
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	ctx        = context.Background()
)

func main() {
	seedHotel("UNCF Hotel", "Japan")
	seedHotel("Yamato Hotel", "Iscandar")
	seedHotel("Andromeda", "Saturn")
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
}

func seedHotel(name, address string) {
	hotel := types.Hotel{
		Name:    name,
		Address: address,
		Rooms:   []primitive.ObjectID{},
	}

	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 99.9,
		},
		{
			Type:      types.DoubleRoomType,
			BasePrice: 129.9,
		},
		{
			Type:      types.DeluxeRoomType,
			BasePrice: 399.9,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
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
