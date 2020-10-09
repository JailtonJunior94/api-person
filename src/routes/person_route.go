package routes

import (
	"github.com/JailtonJunior94/api-person/src/handlers"

	"github.com/gofiber/fiber"
)

// PersonRouter - Rota para CRUD de pessoas
func PersonRouter(app *fiber.App) {
	app.Get("/api/v1/persons", handlers.ListPersonHandle)
	app.Get("/api/v1/persons/:id", handlers.GetByIDPersonHandle)
	app.Post("/api/v1/persons", handlers.NewPersonHandle)
	app.Put("/api/v1/persons/:id", handlers.UpdatePersonHandle)
	app.Delete("/api/v1/persons/:id", handlers.DeletePersonHandle)
}
