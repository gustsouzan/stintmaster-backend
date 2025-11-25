package eventsV1

import (
	"stintmaster/api/domains/events"

	"github.com/gofiber/fiber/v2"
)

func GetEventsHandler(c *fiber.Ctx) error {

	res, err := events.GetEvents()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve events",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
