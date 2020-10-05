package infra

import (
	"context"
	"os"
	"samplegofiber/src/models"
	"time"

	uuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// CreatePerson - Método para criar nova pessoa
func CreatePerson(person models.Person) (*models.Person, error) {
	collection, err := GetCollection(os.Getenv("DbName"), os.Getenv("CollectionName"))
	if err != nil {
		return nil, err
	}

	insert := models.Person{
		ID:        uuid.New().String(),
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Email:     person.Email,
		Age:       person.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = collection.InsertOne(context.Background(), insert)
	if err != nil {
		return nil, err
	}

	return &insert, nil
}

// ListPersons - Método para listagem de pessoas
func ListPersons() ([]models.Person, error) {
	collection, err := GetCollection(os.Getenv("DbName"), os.Getenv("CollectionName"))
	if err != nil {
		return nil, err
	}

	filter := bson.D{{}}
	persons := []models.Person{}

	findResult, findError := collection.Find(context.Background(), filter)
	if findError != nil {
		return persons, findError
	}

	for findResult.Next(context.Background()) {
		p := models.Person{}
		err := findResult.Decode(&p)

		if err != nil {
			return persons, err
		}

		persons = append(persons, p)
	}
	findResult.Close(context.Background())

	if len(persons) == 0 {
		return nil, nil
	}

	return persons, nil
}

// GetPersonByID - Método para listar pessoas através do ID
func GetPersonByID(id string) (*models.Person, error) {
	collection, err := GetCollection(os.Getenv("DbName"), os.Getenv("CollectionName"))
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": id}
	findResult, findError := collection.Find(context.Background(), filter)
	if findError != nil {
		return nil, findError
	}

	person := models.Person{}
	for findResult.Next(context.Background()) {
		p := models.Person{}
		err := findResult.Decode(&p)

		if err != nil {
			return nil, err
		}
		person = p
	}
	findResult.Close(context.Background())

	if len(person.ID) == 0 {
		return nil, nil
	}

	return &person, nil
}

// UpdatePerson - Método para atualizar pessoa
func UpdatePerson(id string, person models.Person) (*models.Person, error) {
	collection, err := GetCollection(os.Getenv("DbName"), os.Getenv("CollectionName"))
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": id}
	update := models.Person{
		ID:        id,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Email:     person.Email,
		Age:       person.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	updateBson := bson.M{"$set": update}

	res, updateError := collection.UpdateOne(context.Background(), filter, updateBson)
	if updateError != nil {
		return nil, updateError
	}

	if res.ModifiedCount == 0 {
		return nil, nil
	}

	return &update, nil
}

// DeletePerson - Método para deletar pessoa
func DeletePerson(id string) (int, error) {
	collection, err := GetCollection(os.Getenv("DbName"), os.Getenv("CollectionName"))
	if err != nil {
		return 0, err
	}

	filter := bson.M{"id": id}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	if res.DeletedCount == 0 {
		return 0, nil
	}

	return 1, nil
}
