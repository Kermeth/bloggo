package repository

import (
	"bloggo/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection = GetDBClient().Database("blog").Collection("posts")

func SavePost(post *model.Post) (*mongo.InsertOneResult, error) {
	return postCollection.InsertOne(context.Background(), post)
}

func GetPosts() ([]model.Post, error) {
	cursor, err := postCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var posts []model.Post
	for cursor.Next(context.Background()) {
		var post model.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
