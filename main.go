package main

import (
	"samplegofiber/src/routes"

	swagger "github.com/arsmn/fiber-swagger"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) {
		c.Send("Seja bem-vindo(a) a API Golang + Fiber + Swagger + MongoDB + Docker")
	})
	routes.PersonRouter(app)

	app.Use("/swagger", swagger.Handler)
	app.Use("/swagger", swagger.New(swagger.Config{
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	println("Running on http://0.0.0.0/3000")
	app.Listen(3000)
}
