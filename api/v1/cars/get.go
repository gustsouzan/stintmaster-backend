package carsV1

import (
	"stintmaster/api/domains/carro"
	"stintmaster/api/domains/pilots"

	"github.com/gofiber/fiber/v2"
)

func GetCars(c *fiber.Ctx) error {

	cars, err := carro.GetAllCars()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(cars)
}

func GetCarClassMap(c *fiber.Ctx) error {

	mappedClasses, err := carro.GetMappedCarClasses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(mappedClasses)
}

func GetCarSuggestions(c *fiber.Ctx) error {
	carSuggestions, err := pilots.GetCarSuggestions()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(carSuggestions)
}

func GetCarClasses(c *fiber.Ctx) error {

	cars, err := carro.GetCarClasses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(cars)
}
