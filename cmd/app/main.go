package main

import (
	"bloggo/internal/auth"
	"bloggo/internal/post"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	server := gin.Default()

	// auth
	auth.Handlers(server)

	// system
	server.GET("/ping", healthcheck)

	// post
	post.Handlers(server)

	err := server.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server")
		return
	}
}

func healthcheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
