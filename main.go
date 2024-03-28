package main

import (
	"github.com/edgartolete/fiber-api-go-starter-template/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func (c *fiber.Ctx)error {
		return c.SendString("Working")
	})

	app.Mount("/users", routes.UsersRoute());

	app.Listen(":7000")
}
