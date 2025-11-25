package pilotsV1

import "github.com/gofiber/fiber/v2"

const (
	HandlerPath = "/pilots"
)

func Register(router fiber.Router) {

	router.Post("/create-pilot", CreatePilotHandler)
	router.Get("", GetPilotsHandler)
}
