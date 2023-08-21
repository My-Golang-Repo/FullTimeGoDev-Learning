package main

import (
	"flag"
	"github.com/fulltimegodev/hotel-reservation-nana/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen Address of API server")
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Andromeda"})
}
