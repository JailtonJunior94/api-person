package infra

import (
	"context"
	"os"
	"time"

	"github.com/JailtonJunior94/api-person/src/config"
	"github.com/JailtonJunior94/api-person/src/models"

	uuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {
	database config.MongoClient
}

func NewUserRepository(database config.MongoClient) Repository {
	return Repository{database: database}
}

func (r Repository) CreatePerson(person models.Person) (*models.Person, error) {
	collection, err := r.database.GetCollection(os.Getenv("DB_NAME"), os.Getenv("COLLECTION_NAME"))
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

func (r Repository) ListPersons() ([]models.Person, error) {
	collection, err := r.database.GetCollection(os.Getenv("DB_NAME"), os.Getenv("COLLECTION_NAME"))
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

func (r Repository) GetPersonByID(id string) (*models.Person, error) {
	collection, err := r.database.GetCollection(os.Getenv("DB_NAME"), os.Getenv("COLLECTION_NAME"))
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

func (r Repository) UpdatePerson(id string, person models.Person) (*models.Person, error) {
	collection, err := r.database.GetCollection(os.Getenv("DB_NAME"), os.Getenv("COLLECTION_NAME"))
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

func (r Repository) DeletePerson(id string) (int, error) {
	collection, err := r.database.GetCollection(os.Getenv("DB_NAME"), os.Getenv("COLLECTION_NAME"))
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
