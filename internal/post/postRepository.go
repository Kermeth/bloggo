package post

import (
	"bloggo/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var db = os.Getenv("MONGO_DB")
var postCollection = repository.GetDBClient().Database(db).Collection("posts")

func SavePost(post *Post) (*mongo.InsertOneResult, error) {
	return postCollection.InsertOne(context.Background(), post)
}

func GetPosts(page, limit int64, search string) ([]Post, error) {
	opts := options.Find().SetSort(bson.D{{"created", 1}}).SetSkip(page * limit).SetLimit(limit)
	var searchOpt bson.D
	if search != "" {
		regex := bson.D{{"$regex", primitive.Regex{Pattern: search, Options: "i"}}}
		searchOpt = bson.D{{"title", regex}}
	} else {
		searchOpt = bson.D{}
	}
	cursor, err := postCollection.Find(context.Background(), searchOpt, opts)
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
