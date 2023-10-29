package main

import (
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
	r.POST("/posts", newPost)
	r.GET("/posts", getPosts)
	//r.GET("/posts/:id", getPost)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server")
		return
	}
}

func newPost(context *gin.Context) {
	var post Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post = createPost(post.Title, post.Content)
	_, err = GetDBClient().Database("blog").Collection("posts").InsertOne(context, post)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post created successfully!", "post": post})
}

func getPosts(context *gin.Context) {
	cursor, err := GetDBClient().Database("blog").Collection("posts").Find(context, bson.D{})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var posts []Post
	for cursor.Next(context) {
		var post Post
		err := cursor.Decode(&post)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, post)
	}
	context.JSON(http.StatusOK, gin.H{"posts": posts})
}

func healthcheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
