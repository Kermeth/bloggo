package auth

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Account struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Created  string `json:"created"`
}

func New(email string, password string) (*Account, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	p := &Account{
		uuid.NewString(),
		email,
		string(hashedPassword),
		time.Now().Format("2006-01-02 15:04:05"),
	}
	return p, nil
}
