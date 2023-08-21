package api

import (
	"github.com/fulltimegodev/hotel-reservation-nana/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Andromeda Yamanami",
		LastName:  "ZZZ-0001-2203",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Aldebaran")
}
