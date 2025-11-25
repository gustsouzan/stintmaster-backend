package healthAPI

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetHealthHandler(c *fiber.Ctx) error {
	c.Context().Response.SetBodyRaw([]byte("OK"))
	c.Context().Response.Header.SetContentType("text/plan")
	return c.SendStatus(http.StatusOK)
}
