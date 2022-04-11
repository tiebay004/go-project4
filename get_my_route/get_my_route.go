package getmyroute

import "github.com/gofiber/fiber/v2"

func Server() {
	app := fiber.New()
	app.Get("/route", func(c *fiber.Ctx) error {
		return c.SendString("This is my route")
	})

	app.Listen(":8889")
}
