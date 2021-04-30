package main

import (
	"log"

	"github.com/JailtonJunior94/api-person/src/routes"

	"github.com/gofiber/fiber"
	"github.com/subosito/gotenv"
)

func main() {
	if err := gotenv.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) {
		c.Send("Seja bem-vindo(a) a API Golang + Fiber + Swagger + MongoDB + Docker + Microsoft Azure")
	})
	routes.PersonRouter(app)

	println("Running on http://0.0.0.0/3000")
	_ = app.Listen(3000)
}
