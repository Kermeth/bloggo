package controller

import (
	"bloggo/config"
	"bloggo/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

var postCollection = config.GetDBClient().Database("blog").Collection("posts")

func CreatePost(context *gin.Context) {
	var post model.Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post = model.New(post.Title, post.Content)
	_, err = postCollection.InsertOne(context, post)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post created successfully!", "post": post})
}

func GetPosts(context *gin.Context) {
	cursor, err := postCollection.Find(context, bson.D{})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var posts []model.Post
	for cursor.Next(context) {
		var post model.Post
		err := cursor.Decode(&post)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, post)
	}
	context.JSON(http.StatusOK, gin.H{"posts": posts})
}
