package infra

import (
	"context"
	"encoding/json"
	"os"
	"samplegofiber/src/models"
)

// CreatePerson método para criar nova pessoa
func CreatePerson(person *models.Person) ([]byte, error) {
	collection, err := GetCollection(os.Getenv("DbName"), os.Getenv("CollectionName"))
	if err != nil {
		return nil, err
	}

	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		return nil, err
	}

	response, _ := json.Marshal(res)
	return response, nil
}
