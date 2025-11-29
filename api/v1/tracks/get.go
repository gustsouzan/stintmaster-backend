package tracksV1

import (
	"stintmaster/api/domains/track"

	"github.com/gofiber/fiber/v2"
)

func GetTracks(c *fiber.Ctx) error {

	tracks, err := track.GetTracks()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(tracks)
}
