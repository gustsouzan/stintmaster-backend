package carsV1

import "github.com/gofiber/fiber/v2"

const (
	HandlerPath = "/cars"
)

func Register(router fiber.Router) {
	router.Get("/", GetCars)
	router.Get("/class-map", GetCarClassMap)
	router.Get("/car-suggestions", GetCarSuggestions)
	router.Get("/classes", GetCarClasses)
}
