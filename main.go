package main

import (
	healthAPI "stintmaster/api/api/health"
	v1 "stintmaster/api/api/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()
	router := app.Group("/api/")

	healthRouter := router.Group(healthAPI.HealthPath)
	healthAPI.Register(healthRouter)
	v1Router := router.Group(v1.HandlerPath)
	v1.Register(v1Router)

	app.Listen(":4040")
}
