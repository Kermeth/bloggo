package auth

import (
	"bloggo/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountsCollection = repository.GetDBClient().Database("blog").Collection("accounts")

func SaveAccount(account *Account) (*mongo.InsertOneResult, error) {
	return accountsCollection.InsertOne(context.Background(), account)
}

func GetAccount(email string) (*Account, error) {
	var account Account
	filter := bson.D{{"email", email}}
	err := accountsCollection.FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
