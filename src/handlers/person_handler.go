package handlers

import (
	"encoding/json"
	"samplegofiber/src/infra"
	"samplegofiber/src/models"

	"github.com/gofiber/fiber"
)

// NewPersonHandle handle para adicionar nova pessoa
func NewPersonHandle(c *fiber.Ctx) {
	var person models.Person
	json.Unmarshal([]byte(c.Body()), &person)

	res, err := infra.CreatePerson(&person)
	if err != nil {
		c.Status(400).Send("Não foi possível cadastrar pessoa")
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}
