package main

import (
	"context"
	"fmt"
	"github.com/fulltimegodev/hotel-reservation-nana/api"
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/fulltimegodev/hotel-reservation-nana/db/fixtures"
	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var (
		mongoDBEndpoint = os.Getenv("MONGO_DB_URL")
		mongoDBName     = os.Getenv("MONGO_DB_NAME")
		ctx             = context.Background()
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDBEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(mongoDBName).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Hotel:   db.NewMongoHotelStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Booking: db.NewMongoBookingStore(client),
	}

	user := fixtures.AddUser(store, "Achilles", "AAA-0004", false)
	fmt.Println("Achilles Token -> ", api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "Admin", "ZZZ-0001-YF", true)
	fmt.Println("Admin Token -> ", api.CreateTokenFromUser(admin))
	hotel := fixtures.AddHotel(store, "Andromeda", "Dolphin Street Andromeda Way", 5, nil)
	room := fixtures.AddRoom(store, "kingsize", true, 99.99, hotel.ID)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2), 2)
	fmt.Println("booking ID ->", booking.ID)

	for i := 0; i < 100; i++ {
		addressLoc := fmt.Sprintf("addres %d", i)
		rating := rand.Intn(5) + 1
		fixtures.AddHotel(store, faker.FirstName(), addressLoc, rating, nil)
	}
	return
}
