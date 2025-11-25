package v1

import (
	eventsV1 "stintmaster/api/api/v1/events"
	pilotsV1 "stintmaster/api/api/v1/pilots"

	"github.com/gofiber/fiber/v2"
)

const (
	HandlerPath = "/v1"
)

func Register(router fiber.Router) {

	eventsRouter := router.Group(eventsV1.HandlerPath)
	pilotsRouter := router.Group(pilotsV1.HandlerPath)
	eventsV1.Register(eventsRouter)
	pilotsV1.Register(pilotsRouter)
}
