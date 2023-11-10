package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handlers(server *gin.Engine) {
	// posts
	server.POST("/signUp", createAccount)
	server.POST("/signIn", login)
}

func createAccount(context *gin.Context) {
	var login *LoginRequest
	err := context.ShouldBindJSON(&login)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := New(login.Email, login.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = SaveAccount(account)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Account created successfully!", "account": account})
}

func login(context *gin.Context) {
	var login *LoginRequest
	err := context.ShouldBindJSON(&login)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := GetAccount(login.Email)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = ComparePasswords(account.Password, login.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ResponseToken(context)
}
