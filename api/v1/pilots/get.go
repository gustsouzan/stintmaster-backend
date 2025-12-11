package pilotsV1

import (
	"stintmaster/api/api/v1/pilots/normalizers"
	dp "stintmaster/api/domains/pilots"

	"github.com/gofiber/fiber/v2"
)

func GetPilots(c *fiber.Ctx) error {

	pilots, err := dp.GetPilot()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve pilots",
		})
	}

	if len(pilots) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No pilots found",
		})
	}

	createdPilots := normalizers.PilotDomainToCreatedPilots(pilots)
	return c.Status(fiber.StatusOK).JSON(createdPilots)
}
