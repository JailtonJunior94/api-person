package handlers

import (
	"encoding/json"

	"github.com/JailtonJunior94/api-person/src/infra"
	"github.com/JailtonJunior94/api-person/src/models"

	"github.com/gofiber/fiber"
)

type Handle struct {
	Repository infra.Repository
}

func NewUserHandle(repository infra.Repository) Handle {
	return Handle{Repository: repository}
}

func (h Handle) ListPersonHandle(c *fiber.Ctx) {
	res, err := h.Repository.ListPersons()
	if err != nil {
		_ = c.Status(400).JSON(models.Response{Success: false, Message: "Não foi possível listar pessoas!"})
		return
	}

	if len(res) == 0 {
		_ = c.Status(404).JSON(models.Response{Success: false, Message: "Não foi encontrado nenhuma pessoa!"})
		return
	}

	_ = c.Status(200).JSON(res)
}

func (h Handle) GetByIDPersonHandle(c *fiber.Ctx) {
	id := c.Params("id")
	res, err := h.Repository.GetPersonByID(id)
	if err != nil {
		_ = c.Status(400).JSON(models.Response{Success: false, Message: "Não foi possível encontrar a pessoa informada"})
		return
	}

	if res == nil {
		_ = c.Status(404).JSON(models.Response{Success: false, Message: "Não foi possível encontrar a pessoa informada"})
		return
	}

	_ = c.Status(200).JSON(res)
}

func (h Handle) NewPersonHandle(c *fiber.Ctx) {
	body := models.Person{}
	_ = json.Unmarshal([]byte(c.Body()), &body)

	newPerson, err := h.Repository.CreatePerson(body)
	if err != nil {
		_ = c.Status(400).JSON(models.Response{Success: false, Message: "Não foi possível cadastrar pessoa"})
		return
	}

	_ = c.Status(201).JSON(newPerson)
}

func (h Handle) UpdatePersonHandle(c *fiber.Ctx) {
	id := c.Params("id")
	body := models.Person{}

	_ = json.Unmarshal([]byte(c.Body()), &body)

	update, err := h.Repository.UpdatePerson(id, body)
	if err != nil {
		_ = c.Status(400).JSON(models.Response{Success: false, Message: "Não foi possível atualizar pessoa"})
		return
	}

	if update == nil {
		_ = c.Status(404).JSON(models.Response{Success: false, Message: "Não foi possível atualizar pessoa"})
		return
	}

	_ = c.Status(200).JSON(update)
}

func (s Handle) DeletePersonHandle(c *fiber.Ctx) {
	id := c.Params("id")
	line, err := s.Repository.DeletePerson(id)
	if err != nil {
		_ = c.Status(400).JSON(models.Response{Success: false, Message: "Erro ao deletar pessoa"})
		return
	}

	if line == 0 {
		_ = c.Status(404).JSON(models.Response{Success: false, Message: "Erro ao deletar pessoa"})
		return
	}

	_ = c.Status(200).JSON(models.Response{Success: true, Message: "Pessoa deletada com sucesso"})
}
