package main

import "github.com/gofiber/fiber/v2"

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to Daily Code!")
}

func main() {
	app := fiber.New()
	app.Get("/", hello)

	app.Listen(":7000")
}
