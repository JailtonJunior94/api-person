package main

import (
	"os"
	"samplegofiber/src/routes"

	"github.com/gofiber/fiber"
	"github.com/subosito/gotenv"
)

func main() {
	initial()

	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) {
		c.Send("Seja bem-vindo(a) a API Golang + Fiber + Swagger + MongoDB + Docker")
	})
	routes.PersonRouter(app)

	println("Running on http://0.0.0.0/3000")
	_ = app.Listen(3000)
}

func initial() {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "Development" {
		_ = gotenv.Load()
	} else {
		_ = gotenv.Load(".env.production")
	}
}
