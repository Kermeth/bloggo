package post

import (
	"bloggo/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var db = os.Getenv("MONGO_DB")
var postCollection = repository.GetDBClient().Database(db).Collection("posts")

func SavePost(post *Post) (*mongo.InsertOneResult, error) {
	return postCollection.InsertOne(context.Background(), post)
}

func GetPosts(page, limit int64) ([]Post, error) {
	opts := options.Find().SetSort(bson.D{{"created", 1}}).SetSkip(page * limit).SetLimit(limit)
	cursor, err := postCollection.Find(context.Background(), bson.D{}, opts)
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
