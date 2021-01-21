package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	Client *mongo.Client
}

func NewConnection() (MongoClient, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("CONNECTION_STRING")))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Printf("Erro ao conectar: %v", err)
		return MongoClient{}, errors.New("Erro ao conectar")
	}

	return MongoClient{Client: client}, nil
}

func (m MongoClient) GetCollection(dbName string, collectionName string) (*mongo.Collection, error) {
	collection := m.Client.Database(dbName).Collection(collectionName)
	return collection, nil
}
