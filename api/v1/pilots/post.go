package pilotsV1

import (
	"stintmaster/api/api/v1/pilots/normalizers"
	"stintmaster/api/domains/pilots"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreatePilot(c *fiber.Ctx) error {

	var req normalizers.PostPilot
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	pilot, err := pilots.CreatePilot(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"pilot": pilot,
	})
}
