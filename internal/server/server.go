package server

import (
	"github.com/gofiber/fiber/v2"
)

func GetHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func PostHandler(c *fiber.Ctx) error {
	return c.SendString("Post request received")
}

func PutHandler(c *fiber.Ctx) error {
	return c.SendString("Put request received")
}

func DeleteHandler(c *fiber.Ctx) error {
	return c.SendString("Delete request received")
}

func SetupRoutes(url string, app *fiber.App) {
	url = "/" + url

	app.Get(url, GetHandler)
	app.Post(url, PostHandler)
	app.Put(url, PutHandler)
	app.Delete(url, DeleteHandler)

	// Additional routes can be added here
}

func NewApp(url string) *fiber.App {
	url = "/" + url
	app := fiber.New()
	// Setup routes
	SetupRoutes(url, app)
	return app
}
