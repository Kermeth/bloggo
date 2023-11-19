package post

import (
	"bloggo/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Handlers(server *gin.Engine) {
	// posts
	server.POST("/posts", auth.CheckToken, createPost)
	server.GET("/posts", getPosts)
}

func createPost(context *gin.Context) {
	var post *Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post = New(post.Title, post.Content)
	_, err = SavePost(post)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post created successfully!", "post": post})
}

func getPosts(context *gin.Context) {
	// Query Params
	page := context.DefaultQuery("page", "0")
	limit := context.DefaultQuery("limit", "10")
	search := context.DefaultQuery("search", "")
	// Convert to int
	pageInt, err := strconv.ParseInt(page, 10, 64)
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Query Posts
	posts, err := GetPosts(pageInt, limitInt, search)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return Posts
	context.JSON(http.StatusOK, gin.H{"posts": posts})
}
