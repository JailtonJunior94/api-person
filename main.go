package main

import (
	"samplegofiber/src/handlers"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Status(200).Send("Eu estou funcionando...")
	})
	app.Post("/person", handlers.NewPersonHandle)

	println("Running on http://0.0.0.0/8080")
	app.Listen(8080)
}
