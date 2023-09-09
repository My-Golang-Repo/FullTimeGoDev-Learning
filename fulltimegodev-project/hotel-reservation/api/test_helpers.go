package api

import (
	"context"
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

type testdb struct {
	client *mongo.Client
	store  *db.Store
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.client.Database(db.DBNAME).Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client)

	return &testdb{
		client: client,
		store: &db.Store{
			User:    db.NewMongoUserStore(client),
			Hotel:   hotelStore,
			Room:    db.NewMongoRoomStore(client, hotelStore),
			Booking: db.NewMongoBookingStore(client),
		},
	}
}
