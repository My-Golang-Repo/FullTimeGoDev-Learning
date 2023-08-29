package main

import (
	"context"
	"flag"
	"github.com/fulltimegodev/hotel-reservation-nana/api"
	"github.com/fulltimegodev/hotel-reservation-nana/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	var (
		hotelStore = db.NewMongoHotelStore(client)
		roomStore  = db.NewMongoRoomStore(client, hotelStore)
		userStore  = db.NewMongoUserStore(client)

		store = &db.Store{
			Hotel: hotelStore,
			Room:  roomStore,
			User:  userStore,
		}

		userHandler  = api.NewUserHandler(userStore)
		hotelHandler = api.NewHotelHandler(store)
		app          = fiber.New(config)
		apiv1        = app.Group("/api/v1")
	)

	listenAddr := flag.String("listenAddr", ":5000", "The listen Address of API server")
	flag.Parse()

	// user handlers
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)

	// hotel handlers
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id/room", hotelHandler.HandleGetHotelRooms)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)

	app.Listen(*listenAddr)
}
