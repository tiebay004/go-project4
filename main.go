package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello , Project 4 test")
	})
	app.Listen(":8889")
	println("test")
}
