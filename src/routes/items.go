package routes

import (
	"skycrypt/src/constants"

	"github.com/gofiber/fiber/v2"
)

func ItemsHandlers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"items": constants.ITEMS,
	})
}
