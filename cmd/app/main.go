package main

import (
	"bloggo/internal/model"
	"bloggo/internal/repository"
	"github.com/gin-gonic/gin"
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

func CreatePost(context *gin.Context) {
	var post *model.Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post = model.New(post.Title, post.Content)
	_, err = repository.SavePost(post)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post created successfully!", "post": post})
}

func GetPosts(context *gin.Context) {
	posts, err := repository.GetPosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"posts": posts})
}
