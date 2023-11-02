package main

import (
	"bloggo/config"
	"bloggo/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func main() {

	r := gin.Default()

	// system
	r.GET("/ping", healthcheck)

	// posts
	r.POST("/posts", CreatePost)
	r.GET("/posts", GetPosts)
	//r.GET("/posts/:id", getPost)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server")
		return
	}
}

func healthcheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}

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
