package database

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "HRM"
const mongoURI = "mongodb://localhost:27017/" + dbName


type MongoInstance struct {
	Client	*mongo.Client
	Db		*mongo.Database 
}

var Mg MongoInstance

func Connect() error {
	ctx, cancel :=context.WithTimeout(context.Background(), 100 * time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	err = client.Connect(ctx)
	db := client.Database(dbName)

	Mg = MongoInstance{
		Client: client,
		Db: db,
	}
	return nil
}