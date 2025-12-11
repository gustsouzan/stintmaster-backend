package pilotsV1

import (
	"log"
	"stintmaster/api/domains/pilots"

	"github.com/gofiber/fiber/v2"
)

func RemovePilot(c *fiber.Ctx) error {

	id := c.Params("id")
	log.Printf("Removing pilot with ID: %s", id)

	err := pilots.RemovePilot(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to remove pilot",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pilot removed successfully",
	})
}
