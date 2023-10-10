package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func (db *DB) Connect() error {
	clientOptions := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))

	var err error
	db.client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Ping() error {
	err := db.client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	return nil
}
