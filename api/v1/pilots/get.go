package pilotsV1

import (
	"stintmaster/api/domains/pilots"

	"github.com/gofiber/fiber/v2"
)

func GetPilotsHandler(c *fiber.Ctx) error {

	res, err := pilots.GetPilots()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve pilots",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
