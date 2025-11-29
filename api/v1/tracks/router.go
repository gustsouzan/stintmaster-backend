package tracksV1

import "github.com/gofiber/fiber/v2"

const (
	HandlerPath = "/tracks"
)

func Register(router fiber.Router) {
	router.Get("/", GetTracks)
}
