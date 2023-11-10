package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

var instance *mongo.Client
var once sync.Once

func GetDBClient() *mongo.Client {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	once.Do(func() {
		mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://"+host+":"+port))
		if err != nil {
			log.Fatal("Unable to connect to MongoDB")
			return
		}
		instance = mongoClient
	})
	return instance
}
