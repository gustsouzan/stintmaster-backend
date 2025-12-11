package pilotsV1

import "github.com/gofiber/fiber/v2"

const (
	HandlerPath = "/pilots"
)

func Register(router fiber.Router) {

	router.Post("/", CreatePilot)
	router.Get("/", GetPilots)
	router.Delete("/:id", RemovePilot)
}
