package infra

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// GetConnection get connection of mongodb
func GetConnection() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("ConnectionString")))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

// GetCollection get collection by dbName and collectionName
func GetCollection(dbName string, collectionName string) (*mongo.Collection, error) {
	client, err := GetConnection()
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return collection, nil
}
