package routes

import (
	"log"

	"github.com/JailtonJunior94/api-person/src/config"
	"github.com/JailtonJunior94/api-person/src/handlers"
	"github.com/JailtonJunior94/api-person/src/infra"

	"github.com/gofiber/fiber"
)

func PersonRouter(app *fiber.App) {
	c, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	r := infra.NewUserRepository(c)
	h := handlers.NewUserHandle(r)

	app.Get("/api/v1/persons", h.ListPersonHandle)
	app.Get("/api/v1/persons/:id", h.GetByIDPersonHandle)
	app.Post("/api/v1/persons", h.NewPersonHandle)
	app.Put("/api/v1/persons/:id", h.UpdatePersonHandle)
	app.Delete("/api/v1/persons/:id", h.DeletePersonHandle)
}
