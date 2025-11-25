package healthAPI

import "github.com/gofiber/fiber/v2"

const (
	HealthPath = "/health"
)

func Register(router fiber.Router) {
	router.Get("", GetHealthHandler)
}
