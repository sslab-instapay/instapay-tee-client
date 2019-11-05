package db

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"log"
	"os"
)

// 몽고 DB 클라이언트 셋업
// https://godoc.org/go.mongodb.org/mongo-driver/mongo
func GetDatabase() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	databaseName := os.Getenv("database_name")
	return client.Database(databaseName), nil
}