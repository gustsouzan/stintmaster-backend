package v1

import (
	carsV1 "stintmaster/api/api/v1/cars"
	eventsV1 "stintmaster/api/api/v1/events"
	pilotsV1 "stintmaster/api/api/v1/pilots"
	tracksV1 "stintmaster/api/api/v1/tracks"

	"github.com/gofiber/fiber/v2"
)

const (
	HandlerPath = "/v1"
)

func Register(router fiber.Router) {

	eventsRouter := router.Group(eventsV1.HandlerPath)
	pilotsRouter := router.Group(pilotsV1.HandlerPath)
	carsRouter := router.Group(carsV1.HandlerPath)
	tracksRouter := router.Group(tracksV1.HandlerPath)
	eventsV1.Register(eventsRouter)
	pilotsV1.Register(pilotsRouter)
	carsV1.Register(carsRouter)
	tracksV1.Register(tracksRouter)
}
