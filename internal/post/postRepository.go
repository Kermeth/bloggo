package post

import (
	"bloggo/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection = repository.GetDBClient().Database("blog").Collection("posts")

func SavePost(post *Post) (*mongo.InsertOneResult, error) {
	return postCollection.InsertOne(context.Background(), post)
}

func GetPosts() ([]Post, error) {
	cursor, err := postCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var posts []Post
	for cursor.Next(context.Background()) {
		var post Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
