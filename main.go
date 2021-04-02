package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nguyenhuuluan434/gofiber/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	setupRoute(app)
	err := app.Listen(":9099")

	if err != nil {
		panic(err)
	}
}

func setupRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})
	api := app.Group("api")
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})
	routes.TodoRoute(api.Group("/todo"))
}
