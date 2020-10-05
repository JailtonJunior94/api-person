package handlers

import (
	"encoding/json"
	"samplegofiber/src/infra"
	"samplegofiber/src/models"

	"github.com/gofiber/fiber"
)

// ListPersonHandle - handle para listar pessoas
func ListPersonHandle(c *fiber.Ctx) {
	res, err := infra.ListPersons()
	if err != nil {
		c.Status(400).JSON(models.Response{Success: false, Message: "Não foi possível listar pessoas!"})
		return
	}

	if len(res) == 0 {
		c.Status(404).JSON(models.Response{Success: false, Message: "Não foi encontrado nenhuma pessoa!"})
		return
	}

	c.Status(200).JSON(res)
}

// GetByIDPersonHandle - handle para recuperar pessoa por Id
func GetByIDPersonHandle(c *fiber.Ctx) {
	id := c.Params("id")
	res, err := infra.GetPersonByID(id)
	if err != nil {
		c.Status(400).JSON(models.Response{Success: false, Message: "Não foi possível encontrar a pessoa informada"})
		return
	}

	if res == nil {
		c.Status(404).JSON(models.Response{Success: false, Message: "Não foi possível encontrar a pessoa informada"})
		return
	}

	c.Status(200).JSON(res)
}

// NewPersonHandle - handle para adicionar nova pessoa
func NewPersonHandle(c *fiber.Ctx) {
	body := models.Person{}
	json.Unmarshal([]byte(c.Body()), &body)

	newPerson, err := infra.CreatePerson(body)
	if err != nil {
		c.Status(400).JSON(models.Response{Success: false, Message: "Não foi possível cadastrar pessoa"})
		return
	}

	c.Status(201).JSON(newPerson)
}

// UpdatePersonHandle - handle para atualizar dados da pessoa
func UpdatePersonHandle(c *fiber.Ctx) {
	id := c.Params("id")
	body := models.Person{}

	json.Unmarshal([]byte(c.Body()), &body)

	update, err := infra.UpdatePerson(id, body)
	if err != nil {
		c.Status(400).JSON(models.Response{Success: false, Message: "Não foi possível atualizar pessoa"})
		return
	}

	if update == nil {
		c.Status(404).JSON(models.Response{Success: false, Message: "Não foi possível atualizar pessoa"})
		return
	}

	c.Status(200).JSON(update)
}

// DeletePersonHandle - handle para remover cadastro da pessoa
func DeletePersonHandle(c *fiber.Ctx) {
	id := c.Params("id")
	line, err := infra.DeletePerson(id)
	if err != nil {
		c.Status(400).JSON(models.Response{Success: false, Message: "Erro ao deletar pessoa"})
		return
	}

	if line == 0 {
		c.Status(404).JSON(models.Response{Success: false, Message: "Erro ao deletar pessoa"})
		return
	}

	c.Status(200).JSON(models.Response{Success: true, Message: "Pessoa deletada com sucesso"})
}
