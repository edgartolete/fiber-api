package routes

import (
	"github.com/edgartolete/fiber-api-go-starter-template/internal/handlers"
	"github.com/gofiber/fiber/v2"
)



func UsersRoute() *fiber.App {

	app := fiber.New();

	app.Get("/", handlers.GetUsersHandler)

	return app;

}