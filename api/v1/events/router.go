package eventsV1

import (
	"github.com/gofiber/fiber/v2"
)

const (
	HandlerPath = "/events"
)

func Register(router fiber.Router) {

	router.Post("/", CreateEvent)
}
