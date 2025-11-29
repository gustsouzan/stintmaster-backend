package eventsV1

import (
	"log"
	n "stintmaster/api/api/v1/events/normalizers"
	"stintmaster/api/domains/events"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {

	var req n.PostEvent
	if err := c.BodyParser(&req); err != nil {
		log.Println("Body parse error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&req); err != nil {
		log.Println("Validation error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	event, err := events.CreateEvent(req)

	if err != nil {
		log.Println("Error creating event:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"event": event,
	})
}
