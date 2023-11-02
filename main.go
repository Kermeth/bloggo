package main

import (
	"bloggo/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	r := gin.Default()

	// system
	r.GET("/ping", healthcheck)

	// posts
	r.POST("/posts", controller.CreatePost)
	r.GET("/posts", controller.GetPosts)
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
